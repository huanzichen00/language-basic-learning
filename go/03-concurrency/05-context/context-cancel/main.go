package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) error {
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("working...", i)
			time.Sleep(300 * time.Millisecond)
		}
	}

	fmt.Println("word completed")
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	err := doWork(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("program finished")
}
