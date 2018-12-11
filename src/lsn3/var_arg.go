package main

import "fmt"

func add(a int, arg ...int) (ret int) {
	ret = a
	for i := 0; i < len(arg); i++ {
		ret += arg[i]
	}
	return
}

func addStr(a string, arg ...string) (ret string) {
	ret = a
	for i := 0; i < len(arg); i++ {
		ret += arg[i]
	}
	return
}

func main() {
	fmt.Println(add(2, 3, 4, 5, 6))
	fmt.Println(addStr("hello", "world", "ok"))
}
