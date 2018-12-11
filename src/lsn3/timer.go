package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Millisecond * 100)
}
func main() {
	nowTime := time.Now()
	fmt.Println(nowTime.Format("2006/01/02 15:04:05"))
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Println("cost %v \n", (end-start)/1000)
}
