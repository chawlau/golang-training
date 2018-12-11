package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func modMap(a map[string]string) {
	a["abc"] = "hellow"
}
func testMap() {
	var a map[string]string
	a = make(map[string]string)
	a["abc"] = "efg"
	a["abd"] = "efg"
	fmt.Println(a)
	modMap(a)
	delete(a, "abd")
	fmt.Println(a)
}

func sortMap() {
	var a map[int]int
	a = make(map[int]int)
	a[1] = 2
	a[3] = 4
	a[2] = 2
	a[4] = 2

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(v, a[v])
	}
}

func main() {
	sortMap()
	fmt.Println("random ", rand.Intn(100))
}
