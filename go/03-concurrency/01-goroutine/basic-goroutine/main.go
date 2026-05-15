package main

import (
	"fmt"
	"time"
)

func say(message string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(message, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go say("goroutine")

	time.Sleep(1 * time.Second)
}
