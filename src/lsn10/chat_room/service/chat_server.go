package service

import (
	"fmt"
	"lsn10/chat_room/client_mgr"
	"lsn10/chat_room/util"
	"net"
)

func RunServer(addr string) (err error) {
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Println(util.GetFileLine() + "tcp listen server failed")
		return
	}

	for {
		var conn net.Conn
		conn, err = listen.Accept()

		if err != nil {
			fmt.Println(util.GetFileLine() + "Accept failed")
			return
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	client := &client_mgr.UserClient{}
	client.SetConn(conn)

	err := client.Process()
	if err != nil {
		fmt.Println(util.GetFileLine()+"client process failed ", err)
		return
	}
}
