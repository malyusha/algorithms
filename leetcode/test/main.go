package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() { run(ctx, 1) }()
	go run(ctx, 2)
	go run(ctx, 3)

	cancel()

	time.Sleep(time.Second * 9)
	fmt.Println("complete")
}

func run(ctx context.Context, num int) {
	<-ctx.Done()

	fmt.Println("run ", num)
}
