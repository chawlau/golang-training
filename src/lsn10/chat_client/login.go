package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"lsn10/chat_room/proto"
	"net"
)

func login(conn net.Conn, userId string, passwd string) (err error) {
	loginReq := &proto.LoginReq{
		UserId: userId,
		Passwd: passwd,
	}

	data, err := json.Marshal(loginReq)
	if err != nil {
		return
	}

	msg := proto.Message{
		Cmd:  proto.UserLogin,
		Data: string(data),
	}

	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data  failed")
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("read package failed, err:", err)
	}

	var loginResp proto.LoginResp
	json.Unmarshal([]byte(msg.Data), &loginResp)

	if loginResp.Err != proto.Errno_SUCCESS {
		fmt.Println("user not register, start register")
		register(conn, userId, passwd)
	} else {
		fmt.Println("login success")
	}

	return
}

func register(conn net.Conn, userId string, passwd string) (err error) {
	req := &proto.RegisterReq{
		User: &proto.UserInfo{
			UserId:   userId,
			Passwd:   passwd,
			NickName: userId,
			Sex:      "male",
			Header:   "http://baidu.com/header/1.jpg",
		},
	}

	data, err := json.Marshal(req)
	if err != nil {
		return
	}

	msg := proto.Message{
		Cmd:  proto.UserRegister,
		Data: string(data),
	}

	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data  failed")
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("read package failed, err:", err)
	}

	var resp proto.RegisterResp

	json.Unmarshal([]byte(msg.Data), &resp)
	if resp.Err != proto.Errno_SUCCESS {
		fmt.Println("user register failed err ", err)
	} else {
		fmt.Println("register success")
	}

	return
}
