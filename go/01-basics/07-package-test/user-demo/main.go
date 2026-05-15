package main

import (
	"fmt"

	"go-basic-learning/01-basics/07-package-test/user-demo/user"
)

func main() {
	repo := user.NewMemoryUserRepository(map[int]user.User{
		1: {
			ID:   1,
			Name: "Tom",
		},
		2: {
			ID:   2,
			Name: "Rose",
		},
	})

	service := user.NewUserService(repo)

	name, err := service.GetUserName(1)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("user name:", name)
	}

	name, err = service.GetUserName(99)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("user name:", name)
	}
}
