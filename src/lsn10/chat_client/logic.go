package main

import (
	"encoding/json"
	"fmt"
	"log"
	"lsn10/chat_room/proto"
	"lsn10/chat_room/util"
	"net"
	"os"
)

func sendTextMessage(conn net.Conn, text string) (err error) {
	sendReq := proto.UserSendMessageReq{
		UserId: *userId,
		Data:   text,
	}

	data, err := json.Marshal(sendReq)

	if err != nil {
		fmt.Println(util.GetFileLine(), "json send req failed")
		return
	}

	sendMsg := proto.Message{
		Cmd:  proto.UserSendMsg,
		Data: string(data),
	}

	data, err = json.Marshal(sendMsg)

	if err != nil {
		fmt.Println(util.GetFileLine(), "json send msg failed")
		return
	}

	return writePackage(conn, data)
}

func listUnReadMsg() {
	select {
	case msg := <-msgChan:
		fmt.Println(msg.UserId, ":", msg.Data)
	default:
		return
	}
}

func enterTalk(conn net.Conn) {
	var msg string
	fmt.Println("please input text")
	fmt.Scanf("%s", &msg)
	sendTextMessage(conn, msg)
}

func enterMenu(conn net.Conn) {
	log.Println("1 list online user")
	log.Println("2 talk")
	log.Println("3 list message")
	log.Println("4 exit")

	var sel int
	fmt.Scanf("%d\n", &sel)

	switch sel {
	case 1:
		outPutUserOnLine()
	case 2:
		enterTalk(conn)
	case 3:
		listUnReadMsg()
	case 4:
		os.Exit(0)
	}
}

func logic(conn net.Conn) {
	enterMenu(conn)
}
