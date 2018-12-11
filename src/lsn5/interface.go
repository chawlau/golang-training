package main

import (
	"fmt"
)

type Task interface {
	Print()
	Sleep()
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

type Test interface {
	Print()
	Sleep()
}

func (p Student) Print() {
	fmt.Println("Print")
}

func (p Student) Sleep() {
	fmt.Println("Sleep")
}

func main() {
	var t Task
	fmt.Println(t)
	var stu Student = Student{
		Name:  "stu1",
		Age:   20,
		Score: 200,
	}

	t = stu
	t.Print()
	t.Sleep()
}
