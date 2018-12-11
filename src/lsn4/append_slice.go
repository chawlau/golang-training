package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]
	s1 := a[2:]
	fmt.Println("%v %v", s, &a[1])

	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	fmt.Println("%v ", s)
	fmt.Println("%v ", s1)
}

func main() {
	testSlice()
}
