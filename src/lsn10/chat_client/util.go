package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"lsn10/chat_room/proto"
	"lsn10/chat_room/util"
	"net"
)

func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte

	n, err := conn.Read(buf[0:4])

	if err != nil || n != 4 {
		fmt.Println("read buf len failed")
		return
	}

	packLen := binary.BigEndian.Uint32(buf[0:4])

	n, err = conn.Read(buf[0:packLen])

	if err != nil || n != int(packLen) {
		fmt.Println("read buf content failed")
		err = errors.New("read body failed")
		return
	}

	err = json.Unmarshal(buf[0:packLen], &msg)

	if err != nil {
		fmt.Println("json unmarshal msg failed")
		return
	}
	return
}

func writePackage(conn net.Conn, data []byte) (err error) {
	var buf [4]byte
	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])

	if err != nil || n != 4 {
		fmt.Println(util.GetFileLine(), "json send package len failed")
		return
	}

	_, err = conn.Write(data)

	if err != nil {
		return
	}

	return
}
