// 表示这是一个可执行程序
package main

// 导入标准库里的格式化输出包
import (
	"errors"
	"fmt"
)

// 定义常量
const language = "Go"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func calculate(a int, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return add(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b)
	default:
		return 0, errors.New("unknown operator")
	}
}

func main() {
	result, err := calculate(10, 2, "+")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 + 2 =", result)
	}

	result, err = calculate(10, 2, "-")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 - 2 =", result)
	}

	result, err = calculate(10, 2, "*")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 * 2 =", result)
	}

	result, err = calculate(10, 0, "/")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}

	result, err = calculate(10, 2, "%")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 % 2 =", result)
	}
}
