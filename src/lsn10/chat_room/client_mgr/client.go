package client_mgr

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"lsn10/chat_room/proto"
	redis_mgr "lsn10/chat_room/redis_mgr"
	"lsn10/chat_room/util"
	"net"

	"github.com/luci/go-render/render"
)

type UserClient struct {
	conn   net.Conn
	userId string
	buf    [8192]byte
}

func (p *UserClient) SetConn(conn net.Conn) {
	p.conn = conn
}

func (p *UserClient) Process() (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			fmt.Println(util.GetFileLine()+"panic", ret)
		}
	}()

	for {
		var msg proto.Message
		msg, err = p.readPackage()

		if err != nil {
			fmt.Println(util.GetFileLine() + "readPackage error")
			ChatClientMgr.DelClient(p.userId)
			p.NotifyOthersUserStatus(p.userId, proto.OFFLINE)
			return
		}

		err = p.processMsg(msg)

		if err != nil {
			fmt.Println(util.GetFileLine()+"processMsg failed err ", err)
			continue
		}
	}
}

func (p *UserClient) readPackage() (msg proto.Message, err error) {
	n, err := p.conn.Read(p.buf[0:4])

	if n != 4 {
		err = errors.New("readPackage failed")
		return
	}

	fmt.Println(util.GetFileLine()+"readPackage ", n)
	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.buf[0:4])
	fmt.Println(util.GetFileLine()+"Package lenth ", packLen)

	n, err = p.conn.Read(p.buf[0:packLen])

	if n != int(packLen) {
		err = errors.New("readPackage body failed")
		return
	}

	fmt.Println(util.GetFileLine()+"recv data ", string(p.buf[0:packLen])+" \n")

	err = json.Unmarshal(p.buf[0:packLen], &msg)

	if err != nil {
		fmt.Println(util.GetFileLine() + "json Unmarshal failed")
	}
	fmt.Println("recv msg ", render.Render(msg))
	return
}

func (p *UserClient) writePackage(data []byte) (err error) {
	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(p.buf[0:4], packLen)

	n, err := p.conn.Write(p.buf[0:4])

	if err != nil {
		fmt.Println(util.GetFileLine() + "write data failed")
		return
	}

	n, err = p.conn.Write(data)

	if err != nil {
		fmt.Println(util.GetFileLine() + "write data failed")
		return errors.New("write data failed")
	}

	if n != int(packLen) {
		fmt.Println(util.GetFileLine() + "write data failed")
		return errors.New("write data not finished")
	}
	return
}

func (p *UserClient) processMsg(msg proto.Message) (err error) {
	switch msg.Cmd {
	case proto.UserLogin:
		err = p.login(msg)
	case proto.UserRegister:
		err = p.register(msg)
	case proto.UserSendMsg:
		err = p.processUserSendMessage(msg)
	default:
		err = errors.New("unsupport message")
		return
	}
	return
}

func (p *UserClient) login(msg proto.Message) (err error) {
	defer func() {
		p.rpcResp("login", err)
	}()

	fmt.Println("recv user login msg data ", msg)
	var req proto.LoginReq
	err = json.Unmarshal([]byte(msg.Data), &req)

	if err != nil {
		fmt.Println(util.GetFileLine() + "json unmarshal loginreq failed")
		return
	}

	_, err = redis_mgr.RedisClientMgr.Login(req.UserId, req.Passwd)

	if err != nil {
		return
	}

	p.userId = req.UserId
	ChatClientMgr.AddClient(p)

	fmt.Println(util.GetFileLine()+"User login succeed", req.UserId)

	p.NotifyOthersUserStatus(req.UserId, proto.ONLINE)
	return
}

func (p *UserClient) NotifyOthersUserStatus(userId string, status proto.UserStatus) {
	users := ChatClientMgr.GetAllUsers()

	users.Range(func(key, value interface{}) bool {
		if key.(string) != userId {
			client := value.(*UserClient)
			client.NotifyUsersStatus(userId, status)
		}
		return true
	})
}

func (p *UserClient) NotifyUsersStatus(userId string, status proto.UserStatus) {
	notify := proto.UserStatusNotifyResp{
		UserId: userId,
		Status: status,
	}

	data, err := json.Marshal(notify)
	if err != nil {
		fmt.Println(util.GetFileLine(), "json NotifyUsersOnLine failed")
		return
	}

	msg := proto.Message{
		Cmd:  proto.UserStatusNotify,
		Data: string(data),
	}

	data, err = json.Marshal(msg)

	if err != nil {
		fmt.Println(util.GetFileLine(), "json NotifyUsersOnLine msg failed")
		return
	}

	err = p.writePackage(data)

	if err != nil {
		fmt.Println(util.GetFileLine(), " writePackage notify failed")
		return
	}
}

func (p *UserClient) register(msg proto.Message) (err error) {
	defer func() {
		p.rpcResp("register", err)
	}()

	fmt.Printf("recv user login msg data %v", msg)
	var req proto.RegisterReq

	err = json.Unmarshal([]byte(msg.Data), &req)

	if err != nil {
		return
	}

	err = redis_mgr.RedisClientMgr.Register(req.User)

	if err != nil {
		return
	}
	ChatClientMgr.AddClient(p)

	p.userId = req.User.UserId

	fmt.Println(util.GetFileLine()+"User register succeed", req.User.UserId)
	return
}

func (p *UserClient) rpcResp(cmd string, err error) {
	resp := &proto.LoginResp{
		Err: proto.Errno_SUCCESS,
	}

	if err != nil {
		switch cmd {
		case "login":
			resp.Err = proto.Errno_LOGIN_FAILED
		case "register":
			resp.Err = proto.Errno_REGISTER_FAILED
		case "send":
			resp.Err = proto.Errno_SEND_FAILED
		}
	}

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(util.GetFileLine() + cmd + "json resp failed")
		return
	}

	msg := &proto.Message{
		Cmd:  proto.UserLoginResp,
		Data: string(data),
	}

	data, err = json.Marshal(msg)

	if err != nil {
		fmt.Println(util.GetFileLine() + cmd + "json respMsg failed")
		return
	}

	err = p.writePackage(data)

	if err != nil {
		fmt.Println(util.GetFileLine() + cmd + "send respPacage failed")
		return
	}
}

func (p *UserClient) processUserSendMessage(msg proto.Message) (err error) {
	defer func() {
		p.rpcResp("send", err)
	}()
	userReq := &proto.UserSendMessageReq{}

	err = json.Unmarshal([]byte(msg.Data), userReq)
	if err != nil {
		fmt.Println(util.GetFileLine(), " unmarshal send Pacage failed")
		return
	}

	users := ChatClientMgr.GetAllUsers()

	users.Range(func(key, value interface{}) bool {
		if key.(string) != userReq.UserId {
			client := value.(*UserClient)
			err = client.SendMessageToUser(userReq.UserId, userReq.Data)
		}
		return true
	})

	return
}

func (p *UserClient) SendMessageToUser(userId string, text string) (err error) {
	recvMsg := proto.UserRecvMessageResp{
		UserId: userId,
		Data:   text,
	}

	data, err := json.Marshal(recvMsg)

	if err != nil {
		fmt.Println(util.GetFileLine(), "json SendMessageToUser failed")
		return
	}

	respMsg := proto.Message{
		Cmd:  proto.UserRecvMsg,
		Data: string(data),
	}

	data, err = json.Marshal(respMsg)

	if err != nil {
		fmt.Println(util.GetFileLine(), "json respMsg failed")
		return
	}

	err = p.writePackage(data)

	if err != nil {
		fmt.Println("send msg write package failed")
	}
	return
}
