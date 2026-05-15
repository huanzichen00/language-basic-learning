package main

import (
	"errors"
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}

func countLeapYears(start int, end int) (int, error) {
	if start > end {
		return 0, errors.New("start year cannot be greater than end year")
	}
	count := 0

	for year := start; year <= end; year++ {
		if isLeapYear(year) {
			count++
		}
	}

	return count, nil
}

func main() {
	start := 2030
	end := 2000

	count, err := countLeapYears(start, end)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("range:", start, "to", end)
	fmt.Println("leap year count:", count)
}
