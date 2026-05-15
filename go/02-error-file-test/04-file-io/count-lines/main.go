package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("open file failed: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("scan file failed: %w", err)
	}

	return count, nil
}

func main() {
	path := "notes.txt"
	content := "first line\nsecond line\nthird line"

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	count, err := countLines(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("line count:", count)
}
