package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("txt.log")
	if err != nil {
		fmt.Println("err ", err)
	}
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')

	defer file.Close()
	if err != nil {
		fmt.Println("read string failed ", err)
		return
	}

	fmt.Println("read str", str)
}
