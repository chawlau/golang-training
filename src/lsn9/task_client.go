package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
	}
}

func MessageSend(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadLine()
		input = string(data)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9097")
	CheckError(err)
	defer conn.Close()

	conn.Write([]byte("hello mingzongyue"))

	fmt.Println("send ok")
}
