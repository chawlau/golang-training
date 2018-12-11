package main

import "fmt"

func add(arg ...int) int {
	a := 0
	for i := 0; i < len(arg); i++ {
		a += arg[i]
	}
	return a
}

var result = func(a int, b int) int {
	return a + b
}

type Slice struct {
	ptr *[100]int
	cap int
	len int
}

func Modify(s Slice, i int) {
	s.ptr[i] = 100
}

func MakeSlice(s Slice, cap int) Slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func ModString(s string) (ret string) {
	s1 := []rune(s)
	s1[0] = 'o'
	ret = string(s1)
	return
}

func ShrinkSlice(slice []int) (s2 []int) {
	if cap(slice) > 0 {
		s2 = slice[:cap(slice)-1]
	}
	return
}

func main() {
	fmt.Println("value ", add(2, 3, 45))
	fmt.Println("value ", result(2, 3))

	//var slice []int
	slice := new([]int)
	*slice = make([]int, 10)
	(*slice)[0] = 1
	(*slice)[1] = 2
	(*slice)[2] = 3
	(*slice)[3] = 4
	(*slice)[4] = 5
	var s2 []int
	s2 = append(s2, (*slice)...)
	fmt.Println("slice", s2)

	var s1 Slice
	s1 = MakeSlice(s1, 100)
	s1.ptr[0] = 24
	Modify(s1, 0)
	fmt.Println("Slice ", s1.ptr)
	fmt.Println(ModString("hloooo"))

	s3 := ShrinkSlice(s2)
	fmt.Println("shrink ", s3)
	s2[0] = 120
	fmt.Println("shrink ", s3)

	var a = 100
	a = a + 1
	fmt.Println("shrink ", a)
}
