package main

import (
	"fmt"
	"lsn2/test"
)

func get_list(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i, " + ", num-i, "=", num)
	}
}

const (
	a = iota
	b
	c
)

func main() {
	get_list(8)
	fmt.Println(test.Name)
}
