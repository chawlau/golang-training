package main

import (
	"fmt"

	go_redis "github.com/astaxie/goredis"
)

func main() {
	var client go_redis.Client
	client.Addr = "127.0.0.1:6379"
	err := client.Set("test", []byte("hello liuchao"))

	if err != nil {
		fmt.Println("err")
	}

	res, err := client.Get("test")
	fmt.Println(string(res))

	f := make(map[string]interface{})

	f["name"] = "zhansgdas"
	f["age"] = 12
	f["sex"] = "male"

	err = client.Hmset("test_hash", f)

	if err != nil {
		fmt.Println("Hmset error")
	}

	_, err = client.Zadd("test_zset", []byte("mingzongyue"), 100)
	if err != nil {
		fmt.Println("Hmset error")
	}
}
