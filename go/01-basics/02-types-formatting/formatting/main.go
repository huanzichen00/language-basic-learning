package main

import "fmt"

const appName = "Go Basic Learning"

func showGlobalInfo() {
	fmt.Println("app name:", appName)
}

func main() {
	showGlobalInfo()

	name := "Tom"
	age := 20
	height := 1.75
	passed := true

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("height:", height)
	fmt.Println("passed:", passed)

	fmt.Printf("name type: %T\n", name)
	fmt.Printf("age type: %T\n", age)
	fmt.Printf("height type: %T\n", height)
	fmt.Printf("passed type: %T\n", passed)

	score := 95
	bonus := 2.5

	finalScore := float64(score) + bonus
	fmt.Println("final score:", finalScore)

	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(message)
}
