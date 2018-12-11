package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	etcd_client "go.etcd.io/etcd/clientv3"
)

const (
	EtcdKey = "/facebook/bacend/logagent/config/10.211.55.6"
)

type CollectConf struct {
	LogPath string `json:"path"`
	Topic   string `json:"topic"`
}

func main() {
	client, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed err ", err)
		return
	}

	fmt.Println("connect succ")
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	var logConfArr []CollectConf
	logConfArr = append(
		logConfArr,
		CollectConf{
			LogPath: "/home/liuchao/Documents/GoCode/logs/" + os.Args[1],
			Topic:   "nginx-log",
		},
	)
	logConfArr = append(
		logConfArr,
		CollectConf{
			LogPath: "/home/liuchao/Documents/GoCode/logs/" + os.Args[2],
			Topic:   "nginx-error",
		},
	)

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed")
		return
	}
	_, err = client.Put(ctx, EtcdKey, string(data))
	cancel()

	if err != nil {
		fmt.Println("put failed err", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := client.Get(ctx, EtcdKey)
	//client.Delete(ctx, EtcdKey)
	cancel()

	if err != nil {
		fmt.Println("get failed err", err)
		return
	}

	for _, val := range resp.Kvs {
		fmt.Println(" key ", string(val.Key), " value ", string(val.Value))
	}
}
