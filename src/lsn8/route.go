package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	for {
		fmt.Println("test")
		time.Sleep(time.Second)
	}
}
func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	fmt.Println(num)
}
