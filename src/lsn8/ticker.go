package main

import (
	"fmt"
	"time"
)

func TickerTest() {
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println("helllo ", v)
	}
}

func SelectTicker() {
	select {
	case <-time.After(time.Second):
		fmt.Println("after")
	}
}

func main() {
	ch := make(chan int, 10)
	chw := make(chan int, 10)

	go func() {
		var i int
		for {
			ch <- i
			time.Sleep(time.Second)
			chw <- i * i
			time.Sleep(time.Second)
			i++
		}
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-chw:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("time out")
		}
	}
}
