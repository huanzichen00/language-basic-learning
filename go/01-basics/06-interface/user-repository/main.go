package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	FindByID(id int) (User, error)
}

type MemoryUserRepository struct {
	users map[int]User
}

func (r MemoryUserRepository) FindByID(id int) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

type UserService struct {
	repo UserRepository
}

func (s UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func main() {
	repo := MemoryUserRepository{
		users: map[int]User{
			1: {
				ID:   1,
				Name: "Tom",
			},
			2: {
				ID:   2,
				Name: "Rose",
			},
		},
	}

	service := UserService{
		repo: repo,
	}

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
