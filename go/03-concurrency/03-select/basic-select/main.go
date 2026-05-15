package main

import (
	"fmt"
	"time"
)

func sendAfter(ch chan string, message string, delay time.Duration) {
	time.Sleep(delay)
	ch <- message
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendAfter(ch1, "from ch1", 500*time.Millisecond)
	go sendAfter(ch2, "from ch2", 1*time.Second)

	for i := 0; i < 2; i++ {
		// 谁先准备好, 就执行谁
		select {
		case msg := <-ch1:
			fmt.Println("received", msg)
		case msg := <-ch2:
			fmt.Println("received", msg)
		}
	}
}
