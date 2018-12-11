package main

func handle(i int, ch chan struct{}) {
	_ = i
	<-ch
}

func main() {
	var maxRoutine = 10
	ch := make(chan struct{}, maxRoutine)
	total := 100

	for i := 0; i < total; i++ {
		ch <- struct{}{}
		go handle(i, ch)
	}
}
