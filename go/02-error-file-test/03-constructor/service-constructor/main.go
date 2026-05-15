package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID   int
	Name string
}

type Repository interface {
	FindByID(id int) (User, error)
}

type MemoryRepository struct {
	users map[int]User
}

func NewMemoryRepository(users map[int]User) *MemoryRepository {
	if users == nil {
		users = map[int]User{}
	}

	return &MemoryRepository{
		users: users,
	}
}

func (r *MemoryRepository) FindByID(id int) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("repo cannot be nil")
	}

	return &Service{
		repo: repo,
	}, nil
}

func (s *Service) GetUserName(id int) (string, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "", fmt.Errorf("get user name failed:, %w", err)
	}
	return user.Name, nil
}

func main() {
	repo := NewMemoryRepository(map[int]User{
		1: {
			ID:   1,
			Name: "Tom",
		},
	})

	service, err := NewService(repo)
	if err != nil {
		fmt.Println("errors:", err)
		return
	}

	name, err := service.GetUserName(1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("user name:", name)

	_, err = NewService(nil)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}
