package main

import (
	"fmt"

	"lsn11/log-agent/conf"
	"lsn11/log-agent/etcd"
	"lsn11/log-agent/kafka"
	"lsn11/log-agent/tailf"

	"github.com/astaxie/beego/logs"
	"github.com/luci/go-render/render"
)

func main() {
	err := conf.LoadConf("ini", "./log_agent.conf")
	if err != nil {
		fmt.Println("load conf failed err ", err)
		panic("load conf failed")
		return
	}

	err = conf.InitLogger()
	if err != nil {
		fmt.Println("init log failed err ", err)
		panic("init log failed")
		return
	}

	logs.Debug("initalsize succ")
	logs.Debug("load conf succ config ", render.Render(conf.AppConfig))

	err = kafka.InitKafka(conf.AppConfig.KafkaAddr)
	if err != nil {
		logs.Error("init kafka ", err)
		return
	}

	collectConf, err := etcd.InitEtcd(conf.AppConfig.EtcdAddr, conf.AppConfig.EtcdKey)
	if err != nil {
		logs.Error("init etcd ", err)
		return
	}

	err = tailf.InitTail(collectConf)
	if err != nil {
		logs.Error("init log failed err ", err)
		return
	}

	logs.Debug("initialize all succ")

	err = serverRun()
	if err != nil {
		logs.Error("server Run err ", err)
		return
	}

	logs.Info("program exited")
	select {}
}
