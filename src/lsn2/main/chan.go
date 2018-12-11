package main

import (
	"fmt"
)

func modify(a chan int) {
	a <- 10
	a <- 12
	a <- 13
	a <- 14
}

func swap(a *int, b *int) {
	var c = *a
	*a = *b
	*b = c
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func init() {
	fmt.Println("init")
}

func main() {
	a := 5
	b := make(chan int, 4)
	c := 6

	modify(b)
	a, c = swap1(a, c)
	fmt.Println("a = ", a)
	fmt.Println("b = ", <-b)
	fmt.Println("b = ", <-b)
	fmt.Println("b = ", <-b)
	fmt.Println("b = ", <-b)
	fmt.Println("c = ", c)
}
