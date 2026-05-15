package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "first"
	ch <- "second"

	fmt.Println(<-ch)

	ch <- "third"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
