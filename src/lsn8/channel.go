package main

import "fmt"

type student struct {
	name string
}

func main() {
	intChan := make(chan interface{}, 1)
	stu := &student{
		name: "stu01",
	}
	intChan <- stu
	stu1 := <-intChan
	stu2, ok := stu1.(*student)
	if ok {
		fmt.Println("ok ", stu2)
	}
}
