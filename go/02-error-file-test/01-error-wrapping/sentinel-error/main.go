package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func findUserName(id int) (string, error) {
	if id != 1 {
		return "", ErrUserNotFound
	}

	return "Tom", nil
}

func getOrderOwnerName(orderID int) (string, error) {
	userID := 99

	name, err := findUserName(userID)
	if err != nil {
		return "", fmt.Errorf("get order owner name failed: %w", err)
	}

	return name, nil
}

func main() {
	name, err := getOrderOwnerName(1001)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("user does not exist")
			return
		}
		fmt.Println("error:", err)
		return
	}
	fmt.Println("owner name:", name)
}
