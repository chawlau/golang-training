package main

import (
	_ "lsn10/chat_room/client_mgr"
	_ "lsn10/chat_room/redis_mgr"
	server "lsn10/chat_room/service"
)

func main() {
	server.RunServer("0.0.0.0:34907")
}
