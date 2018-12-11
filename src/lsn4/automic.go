package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count int32 = 0

func testLock() {
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&count, 1)
			fmt.Println("count ", atomic.LoadInt32(&count))
		}()
	}
}
func main() {
	testLock()
	time.Sleep(10 * time.Second)
}
