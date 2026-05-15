package main

import "fmt"

type SquareResult struct {
	Input  int
	Output int
}

func square(n int, resultCh chan SquareResult) {
	resultCh <- SquareResult{
		Input:  n,
		Output: n * n,
	}
}

func main() {
	numbers := []int{2, 3, 4, 5}

	resultCh := make(chan SquareResult)

	for _, number := range numbers {
		go square(number, resultCh)
	}

	for i := 0; i < len(numbers); i++ {
		result := <-resultCh
		fmt.Println("input:", result.Input, "output:", result.Output)
	}
}
