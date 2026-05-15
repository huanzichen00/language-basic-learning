package main

import (
	"errors"
	"fmt"
)

func findMin(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers cannot be empty")
	}

	min := numbers[0]

	for _, value := range numbers {
		if value < min {
			min = value
		}
	}

	return min, nil
}

func main() {
	testCases := [][]int{
		{13, 7, 42, 9, 21},
		{-3, -10, -1, -8},
		{100},
		{},
	}
	for index, numbers := range testCases {
		min, err := findMin(numbers)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println("case", index, "min:", min)
	}
}
