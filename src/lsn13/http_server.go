package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	server := http.Server{
		Addr: "localhost:9099",
	}

	go func() {
		<-quit
		server.Close()
		server.Shutdown(ctx)
	}()

	go server.ListenAndServe()
	select {
	case <-ctx.Done():
		fmt.Println("quit")
	}
}
