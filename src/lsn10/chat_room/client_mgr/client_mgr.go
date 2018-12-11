package client_mgr

import (
	"errors"
	"fmt"
	"lsn10/chat_room/util"
	"reflect"
	"sync"
)

type ClientMgr struct {
	m         sync.Map
	keyType   reflect.Type
	valueType reflect.Type
}

func NewClientMgr(keyType, valueType reflect.Type) (clientMgr *ClientMgr, err error) {
	if keyType == nil {
		err = errors.New("nil key type")
		return
	}

	if !keyType.Comparable() {
		err = errors.New("key incomparable")
		return
	}

	if valueType == nil {
		err = errors.New("nil value type")
		return
	}

	clientMgr = &ClientMgr{
		keyType:   keyType,
		valueType: valueType,
	}
	return
}

var (
	ChatClientMgr *ClientMgr
)

func init() {
	chatClientMgr, err := NewClientMgr(reflect.TypeOf("string"), reflect.TypeOf(&UserClient{}))
	if err != nil {
		panic(errors.New("init failed"))
	}
	ChatClientMgr = chatClientMgr
}

func (p *ClientMgr) AddClient(client *UserClient) (err error) {
	fmt.Println(util.GetFileLine()+"client ", client.userId, " is online")
	if reflect.TypeOf(client.userId) != p.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(client.userId)))
	}
	if reflect.TypeOf(client) != p.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(client)))
	}

	p.m.Store(client.userId, client)
	err = nil
	return
}

func (p *ClientMgr) DelClient(userId string) (err error) {
	fmt.Println(util.GetFileLine()+"client ", userId, " is offline")
	if reflect.TypeOf(userId) != p.keyType {
		err = errors.New("error key type")
		return
	}
	//delete(p.onLineUsers, userId)
	p.m.Delete(userId)
	return
}

func (p *ClientMgr) GetClient(userId string) (client *UserClient, err error) {
	fmt.Println(util.GetFileLine()+"client ", client.userId, " is offline")
	if reflect.TypeOf(userId) != p.keyType {
		err = errors.New("err get key type")
		return
	}
	val, ok := p.m.Load(userId)
	if !ok {
		err = errors.New("GetClient userId" + userId + " failed")
		return
	}
	client = val.(*UserClient)
	return
}

func (p *ClientMgr) GetAllUsers() *ClientMgr {
	return p
}

func (p *ClientMgr) Range(f func(key, value interface{}) bool) {
	p.m.Range(f)
}
