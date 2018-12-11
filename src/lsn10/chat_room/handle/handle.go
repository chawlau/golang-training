package handle

import (
	"fmt"
	"lsn10/talk_room/client_mgr"
	"lsn10/talk_room/proto"
	_ "lsn10/talk_room/redis_mgr"
)

type Handler struct {
	clientMgr *client_mgr.ClientMgr
}

func NewHandler() *Handler {
	return &Handler {
		clientMgr: NewClientMgr()
	}
}

func (h *Handler) processMessage(msg proto.Message) (err error) {
	switch msg.Cmd {
	case Login:
		fmt.Println(util. GetFileLine() + "Login")
		login(msg.Data)
	case Register:
		fmt.Println(util. GetFileLine() + "Register")
		register(msg.Data)
	}
}

func (h *Handler) login(msg proto.Message) (err error) {
	fmt.Println(util. GetFileLine() + "handle login")
	clientMgr.login(msg)
	return
}

func (h *Handler) register(msg proto.Message) (err error) {
	fmt.Println(util. GetFileLine() + "handle register")
	clientMgr.register(msg)
	return
}
