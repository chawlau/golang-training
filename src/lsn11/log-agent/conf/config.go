package conf

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"
)

var (
	AppConfig *Config
)

type Config struct {
	LogLevel    string
	LogPath     string
	Topic       []string
	ChanSize    int
	KafkaAddr   string
	EtcdAddr    string
	EtcdKey     string
	CollectConf []CollectConf
}

type CollectConf struct {
	LogPath string `json:"path"`
	Topic   string `json:"topic"`
}

func loadCollectConf(conf config.Configer) (err error) {
	path := conf.String("collect::logPath")
	if len(path) == 0 {
		err = errors.New("collect path failed err ")
		return
	}

	topic := conf.String("collect::topic")
	if len(topic) == 0 {
		err = errors.New("collect topic failed err ")
		return
	}

	cc := &CollectConf{
		LogPath: path,
		Topic:   topic,
	}

	AppConfig.CollectConf = append(AppConfig.CollectConf, *cc)

	return
}

func LoadConf(confType, fileName string) (err error) {
	conf, err := config.NewConfig(confType, fileName)

	if err != nil {
		fmt.Println("new config failed err : ", err)
		return
	}

	logLevel := conf.String("logs::logLevel")
	if len(logLevel) == 0 {
		logLevel = "debug"
	}

	logPath := conf.String("logs::logPath")
	if len(logPath) == 0 {
		logPath = "debug"
	}

	chanSize, err := conf.Int("collect::chanSize")
	if err != nil {
		chanSize = 100
	}

	kafkaAddr := conf.String("kafka::addr")
	if len(kafkaAddr) == 0 {
		err = errors.New("init kafka adder failed ")
		return
	}

	etcdAddr := conf.String("etcd::addr")
	if len(etcdAddr) == 0 {
		err = errors.New("init etcd adder failed ")
		return
	}

	etcdKey := conf.String("etcd::etcdKey")
	if len(etcdKey) == 0 {
		err = errors.New("init etcd key failed ")
		return
	}

	AppConfig = &Config{
		LogLevel:  logLevel,
		LogPath:   logPath,
		ChanSize:  chanSize,
		KafkaAddr: kafkaAddr,
		EtcdAddr:  etcdAddr,
		EtcdKey:   etcdKey,
	}

	//interface
	err = loadCollectConf(conf)

	if err != nil {
		fmt.Println("load collect conf failed err ", err)
		return
	}

	return
}
