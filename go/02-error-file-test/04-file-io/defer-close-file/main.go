package main

import "fmt"

// 后进先出
func demoDefer() {
	fmt.Println("start")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("end")
}

func main() {
	demoDefer()
}
