package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextStats struct {
	Lines int
	Words int
	Chars int
}

func countTextStats(path string) (TextStats, error) {
	file, err := os.Open(path)
	if err != nil {
		return TextStats{}, fmt.Errorf("open file failed: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stats := TextStats{}

	for scanner.Scan() {
		line := scanner.Text()

		stats.Lines++
		stats.Words += len(strings.Fields(line))
		stats.Chars += len(line)
	}

	if err := scanner.Err(); err != nil {
		return TextStats{}, fmt.Errorf("scan file failed: %w", err)
	}

	return stats, nil
}

func main() {
	path := "notes.txt"
	content := "hello go\nthis is a test\n"

	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	stats, err := countTextStats(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("lines:", stats.Lines)
	fmt.Println("words:", stats.Words)
	fmt.Println("chars:", stats.Chars)
}
