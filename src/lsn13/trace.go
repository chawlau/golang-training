package main

import (
	"os"
	"runtime/trace"
	"time"
)

func sleep() {
	time.Sleep(5 * time.Second)
	trace.Stop()
}
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	sleep()
}
