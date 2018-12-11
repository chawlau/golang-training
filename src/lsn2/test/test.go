package test

import "fmt"

var Name string = "this is in test package"
var Age int = 1000

func init() {
	fmt.Println("this is a test")
	fmt.Println("this is a test Name", Name)
	fmt.Println("this is a test Age", Age)
	Age = 10
	fmt.Println("this is a test Age", Age)
}
