package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numBytes, err := conn.Read(buf)
		if err != nil {
			continue
		}

		if numBytes != 0 {
			fmt.Println("recive ", string(buf))
		}
	}
}

func main() {
	listenSocket, err := net.Listen("tcp", "127.0.0.1:9097")

	CheckError(err)

	defer listenSocket.Close()

	for {
		conn, err := listenSocket.Accept()
		CheckError(err)
		go MessageSend(conn)
	}
}
