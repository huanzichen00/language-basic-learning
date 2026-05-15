package main

import (
	"errors"
	"fmt"
	"os"
)

func writeTextFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("write file failed: %w", err)
	}

	return nil
}

func readTextFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read file failed: %w", err)
	}

	return string(data), nil
}

func main() {
	path := "hello.txt"
	content := "hello go file\nthis is line two"

	err := writeTextFile(path, content)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	text, err := readTextFile(path)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("file content:")
	fmt.Println(text)

	_, err = readTextFile("missing.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println("error:", err)
			return
		}
	}
	fmt.Println("program finished")
}
