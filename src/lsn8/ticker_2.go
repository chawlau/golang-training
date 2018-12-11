package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		t := time.NewTicker(time.Second)
		select {
		case <-t.C:
			fmt.Println("after")
		}
		//t.Stop()
	}
}
