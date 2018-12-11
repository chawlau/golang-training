package main

import (
	"flag"
	"fmt"
	"log"
	"lsn10/chat_room/proto"
	"net"
)

var (
	flagAddress = flag.String("Address", "0.0.0.0:34907", "address")
	userId      = flag.String("UserId", "00001", "userId")
	passwd      = flag.String("Passwd", "123456", "passwd")
	interf      = flag.String("Interface", "login", "interface")
)
var msgChan chan proto.UserRecvMessageResp

func init() {
	msgChan = make(chan proto.UserRecvMessageResp, 1000)
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	conn, err := net.Dial("tcp", *flagAddress)

	if err != nil {
		fmt.Println("connect server failed err ", err)
		return
	}

	err = login(conn, *userId, *passwd)
	if err != nil {
		fmt.Println("login failed err ", err)
		return
	}

	go processServerMessage(conn)

	for {
		logic(conn)
	}
}
