package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"lsn11/log-agent/conf"
	"lsn11/log-agent/ip"
	"lsn11/log-agent/tailf"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

type EtcdClient struct {
	client *clientv3.Client
	keys   []string
}

var (
	EtcdCli *EtcdClient
)

func InitEtcd(addr, key string) (collectConf []conf.CollectConf, err error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Error("connect failed err ", err)
		return
	}

	logs.Info("connect succ")

	EtcdCli = &EtcdClient{
		client: client,
	}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	for _, ip := range ip.LocalIP {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		EtcdCli.keys = append(EtcdCli.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := EtcdCli.client.Get(ctx, etcdKey)
		if err != nil {
			continue
		}
		cancel()

		logs.Debug("resp from etcd:%v ", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed")
					continue
				}
				fmt.Println("collect conf", collectConf)
			}
		}
	}

	initEtcdWatcher()
	return
}

func initEtcdWatcher() {
	for _, key := range EtcdCli.keys {
		go watchKeys(key)
	}
}

func watchKeys(key string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	for {
		watchResp := cli.Watch(context.Background(), key)
		var collectConf []conf.CollectConf
		getConfSucc := true
		//watchResp是channel,所以阻塞了
		for resp := range watchResp {
			for _, ev := range resp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("etcd delete key[%s]", key)
					continue
				}
				logs.Debug("resp.Events %s %s", ev.Type, ev.Kv.Key)
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Error("key %s Unmarshal failed %s ", key, err)
						getConfSucc = false
						continue
					}
				}

				logs.Debug("get config from etcd %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("UpdateConfig conf", collectConf)
				tailf.UpdateConfig(collectConf)
			}
		}
	}
}
