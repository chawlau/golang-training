package main

import "fmt"

func main() {
	str := "helloworld"
	fmt.Println("len", len([]byte(str)))
	fmt.Println("len", len(str))
}
