package main

import (
	"fmt"
)

func riskyDivide(a int, b int) int {
	if b == 0 {
		panic("division by zero")
	}

	return a / b
}

func safeRun() {
	defer func() {
		value := recover()
		if value != nil {
			fmt.Println("recover from panic:", value)
		}
	}()

	result := riskyDivide(10, 0)
	fmt.Println("result:", result)
}

func unSafeRun() {
	result := riskyDivide(10, 0)
	fmt.Println("result:", result)
}

func main() {
	fmt.Println("start")

	safeRun()

	fmt.Println("end")
}
