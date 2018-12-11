package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	fileName := "/home/liuchao/Documents/GoCode/logs/roll.log"
	tails, err := tail.TailFile(fileName, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		Poll: true,
	})

	if err != nil {
		fmt.Println("tail file err : ", err)
		return
	}

	var msg *tail.Line
	var ok bool

	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, fileName %s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg ", msg)
	}
}
