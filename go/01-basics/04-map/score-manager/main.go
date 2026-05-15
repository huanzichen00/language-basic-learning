package main

import (
	"errors"
	"fmt"
)

func addScore(scores map[string]int, name string, score int) {
	scores[name] = score
}

func findScore(scores map[string]int, name string) (int, error) {
	score, ok := scores[name]
	if !ok {
		return 0, errors.New("student not found")
	}
	return score, nil
}

func updateScore(scores map[string]int, name string, score int) error {
	_, ok := scores[name]
	if !ok {
		return errors.New("student not found")
	}

	scores[name] = score
	return nil
}

func deleteScore(scores map[string]int, name string) {
	delete(scores, name)
}

func averageScore(scores map[string]int) (float64, error) {
	if len(scores) == 0 {
		return 0, errors.New("scores cannot be empty")
	}

	sum := 0

	for _, score := range scores {
		sum += score
	}

	avg := float64(sum) / float64(len(scores))

	return avg, nil
}

func findTopStudent(scores map[string]int) (string, int, error) {
	if len(scores) == 0 {
		return "", 0, errors.New("scores cannot be empty")
	}
	topName := ""
	topScore := 0
	first := true

	// 这里 first 是为了处理第一次赋值。
	// 否则如果所有分数都是负数，topScore := 0 就会出问题。
	for name, score := range scores {
		if first || score > topScore {
			topName = name
			topScore = score
			first = false
		}
	}
	return topName, topScore, nil
}

func main() {
	scores := map[string]int{
		"Tom":  90,
		"Jack": 85,
	}

	addScore(scores, "Rose", 96)
	addScore(scores, "Lucy", 88)

	err := updateScore(scores, "Jack", 91)
	if err != nil {
		fmt.Println("error:", err)
	}

	deleteScore(scores, "Tom")

	score, err := findScore(scores, "Jack")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("Jack score:", score)
	}

	avg, err := averageScore(scores)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("average score: %.2f\n", avg)
	}

	fmt.Println("all scores:", scores)

	stu, score, err := findTopStudent(scores)
	fmt.Println("top student:", stu, "score:", score)
}
