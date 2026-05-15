package main

import (
	"errors"
	"fmt"
)

func average(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers cannot be empty")
	}

	sum := 0

	for _, value := range numbers {
		sum += value
	}

	avg := float64(sum) / float64(len(numbers))

	return avg, nil
}

func main() {
	testCases := [][]int{
		{13, 7, 42, 9, 21},
		{-3, -10, -1, -8},
		{100},
		{},
	}

	for index, numbers := range testCases {
		avg, err := average(numbers)
		if err != nil {
			fmt.Println("case", index, "error:", err)
			continue
		}

		fmt.Printf("case %d average: %.2f\n", index, avg)
	}
}
