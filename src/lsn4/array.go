package main

import "fmt"

func test2() {
	var a [10]int
	b := a
	b[0] = 200
	fmt.Println(a)
}

func testArray() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{1, 2, 3, 4, 5}
	var a3 = [...]int{1: 100, 3: 200}
	fmt.Println(a, a1, a2, a3)
}

func modify(a *[10]int) {
	a[0] = 100
}
func main() {
	var a [10]int
	fmt.Println(a)
	//modify(&a)
	testArray()
}
