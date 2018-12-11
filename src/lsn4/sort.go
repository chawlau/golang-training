package main

import (
	"fmt"
	"sort"
)

func testSort() {
	var a = [...]int{1, 8, 2, 23, 4}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testStrings() {
	var a = [...]string{"abc", "efg", "b", "A"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func main() {
	testStrings()
}
