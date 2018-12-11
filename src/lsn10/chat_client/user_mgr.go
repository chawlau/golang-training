package main

import (
	"fmt"
	"lsn10/chat_room/proto"
)

var onlineUserMap map[string]*proto.UserInfo = make(map[string]*proto.UserInfo, 16)

func outPutUserOnLine() {
	fmt.Println("OnLine user list")

	for id, val := range onlineUserMap {
		if id == *userId {
			continue
		}
		switch val.Status {
		case proto.ONLINE:
			fmt.Println("user ", id, "is on line")
		case proto.OFFLINE:
			fmt.Println("user ", id, "is off line")
		}
	}
}
