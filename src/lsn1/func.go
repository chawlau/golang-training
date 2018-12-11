package main

import (
	"errors"
	"fmt"
	"strings"
)

func ModArray(array [3][]string) {
	array[0][1] = "ok"
}

type Operate func(a, b int) int

type CalFunc func(x, y int) (int, error)

func GetFunc(op Operate) CalFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("Nil")
		}
		return op(x, y), nil
	}
}

func MakeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	ModArray(complexArray1)
	fmt.Println(complexArray1)

	func1 := MakeSuffixFunc(".jpg")
	func2 := MakeSuffixFunc(".png")
	fmt.Println(func1("test"))
	fmt.Println(func2("test.png"))
}
