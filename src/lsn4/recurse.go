package main

import (
	"fmt"
	"time"
)

func recusive(n int) {
	fmt.Println("hello")
	time.Sleep(time.Second)

	if n > 10 {
		return
	}
	recusive(n + 1)
}

func main() {
	recusive(0)
}
