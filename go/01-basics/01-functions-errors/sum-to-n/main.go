package main

import (
	"errors"
	"fmt"
)

func sumToN(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum, nil
}

func main() {
	result, err := sumToN(100)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("sum from 1 to 100:", result)

	result, err = sumToN(0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("sum from 1 to 0:", result)
}
