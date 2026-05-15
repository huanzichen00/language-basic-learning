package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func printUserName(user *User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	fmt.Println("user name:", user.Name)
	return nil
}

func main() {
	user := &User{
		ID:   1,
		Name: "Tom",
	}

	err := printUserName(user)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var missingUser *User

	err = printUserName(missingUser)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
