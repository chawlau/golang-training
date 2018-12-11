package main

import "fmt"

func add() {
	i := 0
	p := &i
	defer fmt.Println("defer i ", *p)
	defer fmt.Println("defer second ")
	i++
}

func main() {
	add()
}
