package main

import "fmt"

func main() {
	// nil slice == nil
	var numbers []int
	// empty slice == []
	// 后端常用这个
	emptyNumbers := []int{}

	fmt.Println("numbers:", numbers)
	fmt.Println("emptyNumbers:", emptyNumbers)

	fmt.Println("numbers == nil:", numbers == nil)
	fmt.Println("emptyNubmers == nil", emptyNumbers == nil)

	fmt.Println("len(numbers):", len(numbers))
	fmt.Println("len(emptyNumbers):", len(emptyNumbers))

	numbers = append(numbers, 10)
	emptyNumbers = append(emptyNumbers, 20)

	fmt.Println("numbers after append:", numbers)
	fmt.Println("emptyNumbers after append:", emptyNumbers)

	// map 类似 slice 但是:
	// nil map 不能写入
	var scores map[string]int
	// empty map 可以写入
	emptyScores := map[string]int{}
	// 都可以读取不存在的 key

	fmt.Println("scores == nil:", scores == nil)
	fmt.Println("emptyScores == nil:", emptyScores == nil)

	fmt.Println("len(scores):", len(scores))
	fmt.Println("len(emptyScores):", len(emptyScores))

	emptyScores["Tom"] = 90
	fmt.Println("emptyScores:", emptyScores)

	scores["Tom"] = 90
}
