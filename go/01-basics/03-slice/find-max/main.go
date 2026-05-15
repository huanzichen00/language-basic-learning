package main

import (
	"errors"
	"fmt"
)

func findMax(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers cannot be empty")
	}

	max := numbers[0]

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	return max, nil
}

func main() {
	testCases := [][]int{
		{13, 7, 42, 9, 21},
		{-3, -10, -1, -8},
		{100},
		{},
	}

	for index, numbers := range testCases {
		max, err := findMax(numbers)
		if err != nil {
			fmt.Println("case", index, "error:", err)
			continue
		}

		fmt.Println("case", index, "max", max)
	}
}
