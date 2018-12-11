package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOPATH")
	fmt.Println("GOOS ", goos)
}
