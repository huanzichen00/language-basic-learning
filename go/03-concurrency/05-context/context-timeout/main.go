package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("work completed")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := doWork(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("program finished")
}
