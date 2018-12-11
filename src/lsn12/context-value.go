package main

import (
	"context"
	"fmt"
)

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace-id").(int)
	if !ok {
		ret = 123456
	}
	s, _ := ctx.Value("session").(string)
	fmt.Println("ret ", ret, " session ", s)
}

func main() {
	ctx := context.WithValue(context.Background(), "trace-id", "123456")
	ctx = context.WithValue(context.Background(), "session", "asdsadsadsalasdasdsadsa")
	process(ctx)
}
