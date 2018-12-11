package main

import "fmt"

type Student struct {
	Name string
	Sex  string
}

func Test(a interface{}) {
	b, ok := a.(Student)
	if !ok {
		fmt.Println("Failed")
		return
	}
	fmt.Println(b)
}

func main() {
	var b Student
	Test(b)
}
