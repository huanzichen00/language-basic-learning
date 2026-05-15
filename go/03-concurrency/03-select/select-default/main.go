package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "hello"
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("received:", msg)
			fmt.Println("program finished")
			return
		default:
			fmt.Println("waiting...")
			time.Sleep(300 * time.Millisecond)
		}
	}

	fmt.Println("program finished")
}
