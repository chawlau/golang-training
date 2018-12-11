package main

import "fmt"

type slice struct {
	ptr *[100]int
	len int
	cap int
}

func makeSlice(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func testSlice2() {
	var s1 slice
	s1 = makeSlice(s1, 10)
	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)
}

func testSlice3() {
	var b []int = []int{1, 2, 3, 4, 5}
	fmt.Println(b)
	modify1(b)
	fmt.Println(b)
}

func main() {
	testSlice2()
}
