package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return "path " + p.path + " op " + p.op + " createTime " + p.createTime + " message " + p.message
}

func Open(fileName string) error {
	file, err := os.Open("test.ssdsd")

	if err != nil {
		return &PathError{
			path:       "test.ssdsd",
			op:         "read",
			createTime: fmt.Sprintf("%v", time.Now()),
			message:    err.Error(),
		}
	}

	defer file.Close()
	return nil
}

func main() {
	err := Open("asdsadsa.txt")
	v, ok := err.(*PathError)

	if ok {
		fmt.Println("get path error", v)
	}

	switch err.(type) {
	case *PathError:
		fmt.Println("get path error ", v)
	default:
	}
}
