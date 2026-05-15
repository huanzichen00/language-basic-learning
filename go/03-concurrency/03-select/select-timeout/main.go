package main

import (
	"fmt"
	"time"
)

func slowTask(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "task finished"
}

func main() {
	ch := make(chan string)

	go slowTask(ch)

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}
