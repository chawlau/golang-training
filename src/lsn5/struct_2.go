package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type integer int

type Cart struct {
	name string
	age  int
}

func (t *Cart) Run() {
	fmt.Println("hello worlld")
}

type Train struct {
	Cart
	start time.Time
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (p *Student) init(name string, age int, score int) {
	p.Name = name
	p.Age = age
	p.Score = score
}

func (p Student) get() Student {
	return p
}

func main() {
	var i integer = 1000
	var j int = 20
	fmt.Println(i, j)

	/*
		stu := Student{
			Name:  "stu01",
			Age:   18,
			Score: 80,
		}
	*/
	var stu Student
	stu.init("stu", 10, 200)

	data, err := json.Marshal(stu)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(string(data))

	var t Train
	t.Cart.name = "helldasd"
	t.age = 32
	fmt.Println("t", t)
	t.Run()
	stu1 := stu.get()
	fmt.Println("t", stu1)
}
