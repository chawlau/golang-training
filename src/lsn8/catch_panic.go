package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic ", err)
		}
	}()
	var m map[string]int
	m["stu"] = 3
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	go test()
	for i := 0; i < 1024; i++ {
		fmt.Println("work ", i)
	}

	time.Sleep(time.Second * 100)
}
