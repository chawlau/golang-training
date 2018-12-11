package main

import "fmt"

func changeString() {
	s := "我  hello world"
	fmt.Println(s)
	s1 := []rune(s)
	s1[1] = '0'

	str := string(s1)
	fmt.Println(str)
}

func main() {
	changeString()
}
