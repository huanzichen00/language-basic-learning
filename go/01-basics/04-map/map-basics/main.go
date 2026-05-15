package main

import "fmt"

func main() {
	scores := map[string]int{
		"Tom":  90,
		"Jack": 85,
		"Rose": 96,
	}

	fmt.Println("scores:", scores)

	scores["Lucy"] = 88
	fmt.Println("after add Lucy:", scores)

	scores["Jack"] = 91
	fmt.Println("after add Jack:", scores)

	delete(scores, "Tom")
	fmt.Println("after delete Tom:", scores)

	roseScore := scores["Rose"]
	fmt.Println("Rose score:", roseScore)

	value, ok := scores["Tom"]
	if ok {
		fmt.Println("Tom score:", value)
	} else {
		fmt.Println("Tom not found")
	}
	for name, score := range scores {
		fmt.Println(name, score)
	}
}
