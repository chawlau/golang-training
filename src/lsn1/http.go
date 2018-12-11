package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	for i := 0.0; i < 100000000; i++ {
		math.Sqrt(i)
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	//http.HandleFunc("/", sayhelloName)          //设置访问的路由
	err := http.ListenAndServe(os.Args[1], nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
