package ip

import (
	"fmt"
	"net"

	"github.com/astaxie/beego/logs"
)

var (
	LocalIP []string
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip failed %v", err))
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				LocalIP = append(LocalIP, string(ipnet.IP.String()))
			}
		}
	}
	logs.Info("LocalIP ", LocalIP)
}
