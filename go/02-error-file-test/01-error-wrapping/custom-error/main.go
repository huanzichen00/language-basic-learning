package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ":" + e.Message
}

func createUser(name string, age int) error {
	if name == "" {
		return ValidationError{
			Field:   "name",
			Message: "cannot be empty",
		}
	}

	if age <= 0 {
		return ValidationError{
			Field:   "age",
			Message: "must be greater than 0",
		}
	}
	return nil
}

func registerUser(name string, age int) error {
	err := createUser(name, age)
	if err != nil {
		return fmt.Errorf("register user failed: %w", err)
	}

	return nil
}

func main() {
	err := registerUser("", 20)
	if err != nil {
		var validationErr ValidationError

		if errors.As(err, &validationErr) {
			fmt.Println("validation field:", validationErr.Field)
			fmt.Println("validation message:", validationErr.Message)
			return
		}

		fmt.Println("error:", err)
		return
	}
	fmt.Println("register success")
}
