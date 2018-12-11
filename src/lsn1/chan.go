package main

import (
	"fmt"
	"time"
)

func SelectSkip() {
	intChan := make(chan int, 1)
	intChan2 := make(chan int, 1)
	intChan3 := make(chan int, 1)
	intChan2 <- 2
	close(intChan)
	for {
		select {
		case val, ok := <-intChan:
			if ok {
				fmt.Println("val ", val)
			} else {
				fmt.Println("change chan to nil")
				intChan = nil
			}
		case val, ok := <-intChan2:
			if ok {
				fmt.Println("val", val)
			} else {
				fmt.Println("no val")
			}
		case <-intChan3:
			fmt.Println("case 3")
			intChan3 = nil
		default:
			fmt.Println("continue")
		}
		time.Sleep(1 * time.Second)
	}
}

func SetValue(intChan chan int, sleep int) {
	time.Sleep(time.Duration(sleep) * time.Second)
	intChan <- sleep
}

func BreakForSelect() {
	intChan := make(chan int, 1)
	intChan2 := make(chan int, 1)
	intChan3 := make(chan int, 1)
	countChan := 0
	go SetValue(intChan, 1)
	go SetValue(intChan2, 2)
	go SetValue(intChan3, 3)

	for {
		select {
		case val, ok := <-intChan:
			if ok {
				fmt.Println("val ", val)
				countChan++
				close(intChan)
				intChan = nil
			}

		case val, ok := <-intChan2:
			if ok {
				fmt.Println("val ", val)
				countChan++
				close(intChan2)
				intChan2 = nil
			}

		case val, ok := <-intChan3:
			if ok {
				fmt.Println("val ", val)
				countChan++
				close(intChan3)
				intChan3 = nil
			}

		default:
			fmt.Println("dealult")
			if countChan == 3 {
				goto end
			}
			time.Sleep(1 * time.Second)
		}
	}
end:
	fmt.Println("End")
}

func BreakForSelect2() {
	intChan := make(chan int, 1)
	intChan2 := make(chan int, 1)
	intChan3 := make(chan int, 1)
	countChan := 0
	go SetValue(intChan, 1)
	go SetValue(intChan2, 2)
	go SetValue(intChan3, 3)

loop:
	for {
		select {
		case val, ok := <-intChan:
			if ok {
				fmt.Println("val ", val)
				countChan++
				close(intChan)
				intChan = nil
			}

		case val, ok := <-intChan2:
			if ok {
				fmt.Println("val ", val)
				countChan++
				close(intChan2)
				intChan2 = nil
			}

		case val, ok := <-intChan3:
			if ok {
				fmt.Println("val ", val)
				countChan++
				close(intChan3)
				intChan3 = nil
			}

		default:
			fmt.Println("dealult")
			if countChan == 3 {
				break loop
			}
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("End")
}

func main() {
	//SelectSkip()
	BreakForSelect2()
}
