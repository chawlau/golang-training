package main

import (
	"fmt"
	"time"
)

const (
	Male   = 1
	Female = 2
)

func main() {
	second := time.Now().Unix()

	for {
		if second%Female == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("man")
		}

		time.Sleep(300 * time.Millisecond)
	}
}
