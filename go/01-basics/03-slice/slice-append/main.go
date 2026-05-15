package main

import "fmt"

// 修改元素会生效
func doubleValues(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 2
	}
}

func appendValue(numbers []int) {
	numbers = append(numbers, 100)
	fmt.Println("inside appendValue:", numbers)
}

func main() {
	numbers := []int{1, 2, 3}

	doubleValues(numbers)
	fmt.Println("after doubleValues:", numbers)

	appendValue(numbers)
	fmt.Println("after appendValue:", numbers)

	numbers = append(numbers, 100)
	fmt.Println("after append in main", numbers)
}
