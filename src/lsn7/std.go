package main

import (
	"fmt"
	"os"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	file, _ := os.OpenFile("txt.log", os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Fprintf(file, "do balance err")
	var str = "huas 18 20.34"
	var stu student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}
