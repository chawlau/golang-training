package main

import (
	"fmt"
	"time"
)

func test1(boolChan chan bool) {
	time.Sleep(time.Second * 1)
	fmt.Println("test1")
	boolChan <- true
}

func test2(boolChan chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("test2")
	boolChan <- true
}
func test3(boolChan chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println("test3")
	boolChan <- true
}

func main() {
	boolChan := make(chan bool, 3)

	go test1(boolChan)
	go test2(boolChan)
	go test3(boolChan)
	for i := 0; i < 3; i++ {
		<-boolChan
	}
}
