package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "./log_agent.conf")

	if err != nil {
		fmt.Println("new config failed err : ", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server:port failed, err ", err)
		return
	}

	fmt.Println("Port: ", port)
	logLevel := conf.String("logs::logLevel")

	if len(logLevel) == 0 {
		logLevel = "debug"
	}

	fmt.Println("logLevel: ", logLevel)

	logPath := conf.String("logs::logPath")
	fmt.Println("logPath: ", logPath)
}
