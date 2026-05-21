package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "stopped:", ctx.Err())
			return
		default:
			fmt.Println("worker", id, "working")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(1 * time.Second)

	cancel()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("main finished")
}
