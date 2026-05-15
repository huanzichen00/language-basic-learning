package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 10
	ch <- 20
	close(ch)

	value, ok := <-ch
	fmt.Println("first:", value, ok)

	value, ok = <-ch
	fmt.Println("second:", value, ok)

	value, ok = <-ch
	fmt.Println("third:", value, ok)
}
