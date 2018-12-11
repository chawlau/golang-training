package main

import (
	"encoding/json"
	"fmt"
	"lsn10/chat_room/proto"
	"lsn10/chat_room/util"
	"net"
	"os"
)

func processServerMessage(conn net.Conn) {
	for {
		msg, err := readPackage(conn)

		if err != nil {
			fmt.Println(util.GetFileLine(), "read package failed", err)
			os.Exit(0)
			return
		}

		switch msg.Cmd {
		case proto.UserRecvMsg:
			recvMessageFromServer(msg)
		case proto.UserStatusNotify:
			updateUserOnLineStatus(msg)
		}
	}
}

func recvMessageFromServer(msg proto.Message) {
	var recvMsg proto.UserRecvMessageResp
	err := json.Unmarshal([]byte(msg.Data), &recvMsg)

	if err != nil {
		fmt.Println(util.GetFileLine(), "unmarshal recv msg failed", err)
		return
	}

	fmt.Println("userId", recvMsg.UserId, "msg", recvMsg.Data)
}

func updateUserOnLineStatus(msg proto.Message) {
	notifyMsg := &proto.UserStatusNotifyResp{}
	err := json.Unmarshal([]byte(msg.Data), notifyMsg)

	if err != nil {
		fmt.Println(util.GetFileLine(), "updateUserOnLineStatus json unmarshal failed", err)
		return
	}

	user, ok := onlineUserMap[notifyMsg.UserId]

	if !ok {
		user = &proto.UserInfo{
			UserId: notifyMsg.UserId,
		}
	}

	user.Status = notifyMsg.Status
	onlineUserMap[notifyMsg.UserId] = user
	outPutUserOnLine()
}
