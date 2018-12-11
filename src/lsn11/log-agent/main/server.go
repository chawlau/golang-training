package main

import (
	"time"

	"lsn11/log-agent/kafka"
	"lsn11/log-agent/tailf"

	"github.com/astaxie/beego/logs"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetLineMsg()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed err ", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	//logs.Debug("read msg ", msg.Msg, " topic ", msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
