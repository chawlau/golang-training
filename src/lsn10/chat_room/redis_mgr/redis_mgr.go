package redis_mgr

import (
	"encoding/json"
	"flag"
	"fmt"
	"lsn10/chat_room/proto"
	"lsn10/chat_room/util"
	"time"

	redis "github.com/garyburd/redigo/redis"
)

var (
	UserTable      = "users"
	RedisClientMgr *RedisMgr
)

type RedisMgr struct {
	pool *redis.Pool
}

var (
	redisAddr = flag.String("RedisAddr", "127.0.0.1:6379", "redisAddr")
	idleConn  = flag.Int("IdleConn", 16, "idleConn")
	maxConn   = flag.Int("MaxConn", 1024, "maxConn")
	timeOut   = flag.Int("TimeOut", 300, "timeOut")
)

func init() {
	RedisClientMgr = &RedisMgr{
		pool: &redis.Pool{
			MaxIdle:     *idleConn,
			MaxActive:   *maxConn,
			IdleTimeout: time.Duration(5) * time.Second,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", *redisAddr)
			},
		},
	}
}

func (p *RedisMgr) GetConn() redis.Conn {
	return p.pool.Get()
}

func (p *RedisMgr) PutConn(conn redis.Conn) {
	conn.Close()
}

func (p *RedisMgr) getUser(conn redis.Conn, userId string) (user *proto.UserInfo, err error) {
	res, err := redis.String(conn.Do("Hget", UserTable, userId))

	//star_coding
	//include key not exist and user not exsit
	if err != nil {
		fmt.Println("get User err ", err)
		if err == redis.ErrNil {
			err = proto.ErrUserNotExist
		}
		return
	}

	user = &proto.UserInfo{}
	err = json.Unmarshal([]byte(res), user)

	if err != nil {
		return
	}
	return
}

func (p *RedisMgr) Login(userId string, passwd string) (user *proto.UserInfo, err error) {
	fmt.Println(util.GetFileLine() + " redis handle login")

	conn := p.pool.Get()
	defer conn.Close()

	user, err = p.getUser(conn, userId)

	if err != nil {
		return
	}

	if user.UserId != userId || user.Passwd != passwd {
		err = proto.ErrInvalidPasswd
		return
	}

	user.Status = proto.ONLINE
	user.LastLogin = util.GetTodayDateTime()
	return
}

func (p *RedisMgr) Register(user *proto.UserInfo) (err error) {
	fmt.Println(util.GetFileLine() + " redis handle register")
	conn := p.pool.Get()
	defer conn.Close()

	if user == nil {
		fmt.Println(util.GetFileLine() + "invalid user")
		err = proto.ErrInvalidParams
		return
	}

	_, err = p.getUser(conn, user.UserId)

	//has exist
	if err == nil {
		err = proto.ErrUserExist
		return
	}

	//not unregister
	if err != proto.ErrUserNotExist {
		return
	}

	data, err := json.Marshal(user)

	if err != nil {
		return
	}
	_, err = conn.Do("HSet", UserTable, user.UserId, string(data))

	if err != nil {
		return
	}
	return
}
