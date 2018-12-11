package main

import (
	"fmt"
)

func reverse(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(tmp)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}

func main() {
	var str = "hello world \n"
	var str1 = `
	hello \t \n
	world
	`
	var b byte = 'c'
	var c bool = true
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Printf("%T \n", b)
	fmt.Printf("%t \n", c)
	fmt.Println("res", reverse(str))
}
