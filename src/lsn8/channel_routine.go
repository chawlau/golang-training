package main

import (
	"fmt"
)

func read(ch chan int) {
	for {
		val := <-ch
		fmt.Println("val ", val)
	}
}

func calc(taskChan chan int, resChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resChan <- v
		}
	}

	fmt.Println("exit")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	for i := 0; i < 100; i++ {
		intChan <- i
	}

	close(intChan)

	for i := 0; i < 8; i++ {
		go calc(intChan, resChan, exitChan)
	}

	//等待所有计算的goroutine全部退出
	for i := 1; i < 8; i++ {
		<-exitChan
	}

	close(resChan)
	for v := range resChan {
		fmt.Println("val ", v)
	}
}
