package main

import "fmt"

func printNumbers(numbers []int) {
	fmt.Println("numbers:", numbers)
	fmt.Println("len:", len(numbers))
	fmt.Println("cap:", cap(numbers))
}

func main() {
	numbers := []int{10, 20, 30}

	printNumbers(numbers)

	numbers = append(numbers, 40)
	numbers = append(numbers, 50)

	printNumbers(numbers)

	part := numbers[1:4]
	fmt.Println("part:", part)

	for index, value := range numbers {
		fmt.Println("index:", index, "value:", value)
	}
}
