package main

import "errors"

var ErrUserNotFound = errors.New("user not found")

type UserService struct {
	repo InMemoryUserRepository
}

func (s UserService) ListUsers() []User {
	return s.repo.List()
}

func (s UserService) GetUserByID(id int) (User, error) {
	user, ok := s.repo.GetByID(id)
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

func (s UserService) CreateUser(name string) User {
	return s.repo.Create(name)
}

func (s UserService) UpdateUser(id int, name string) (User, error) {
	user, ok := s.repo.Update(id, name)
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

func (s UserService) DeleteUser(id int) error {
	ok := s.repo.Delete(id)
	if !ok {
		return ErrUserNotFound
	}

	return nil
}
