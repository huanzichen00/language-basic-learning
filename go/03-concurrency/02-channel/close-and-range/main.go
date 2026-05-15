package main

import "fmt"

func produceNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	close(ch)

	ch <- 100
}

func main() {
	ch := make(chan int)

	go produceNumbers(ch)

	for value := range ch {
		fmt.Println("received:", value)
	}

	fmt.Println("channel closed")
}
