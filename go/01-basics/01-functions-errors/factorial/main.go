package main

import (
	"errors"
	"fmt"
)

func factorial(n int) (int, error) {
	if n < 0 {
		err := errors.New("n cannot be negative")
		return 0, err
	}

	if n == 0 {
		return 1, nil
	}

	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res, nil
}

func cal(n int) {
	num, err := factorial(n)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%d! = %d\n", n, num)
}

func main() {
	cal(5)
	cal(0)
	cal(-1)
}
