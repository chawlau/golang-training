package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	for {
		select {
		case v := <-ch:
			fmt.Println("val ", v)
		default:
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
			goto end
		}
	}
end:
}
