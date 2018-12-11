package main

import "fmt"

func main() {
	j := new(int)
	*j = 100
	fmt.Println(*j)
}
