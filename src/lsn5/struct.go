package main

import "fmt"

type Student struct {
	Name  string
	Age   int64
	score float32
}

func main() {
	var stu Student
	stu.Age = 18
	stu.Name = "ming1234523242342432432"
	stu.score = 88.00

	fmt.Println("stu %p ", &stu.Name)
	fmt.Println("stu %p ", &stu.Age)
	fmt.Println("stu %p ", &stu.score)
}
