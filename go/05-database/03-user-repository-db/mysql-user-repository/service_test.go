package main

import (
	"errors"
	"testing"
)

type fakeUserRepository struct {
	listUsersResult []User
	listUsersErr    error

	getUserResult     User
	getUserErr        error
	getUserCalledWith int64

	createUserResult     User
	createUserErr        error
	createUserCalledWith string

	updateUserResult     User
	updateUserErr        error
	updateUserCalledID   int64
	updateUserCalledName string

	deleteUserErr        error
	deleteUserCalledWith int64
}

func (f *fakeUserRepository) ListUsers() ([]User, error) {
	return f.listUsersResult, f.listUsersErr
}

func (f *fakeUserRepository) GetUserByID(id int64) (User, error) {
	f.getUserCalledWith = id
	return f.getUserResult, f.getUserErr
}

func (f *fakeUserRepository) CreateUser(name string) (User, error) {
	f.createUserCalledWith = name
	return f.createUserResult, f.createUserErr
}

func (f *fakeUserRepository) UpdateUser(id int64, name string) (User, error) {
	f.updateUserCalledID = id
	f.updateUserCalledName = name
	return f.updateUserResult, f.updateUserErr
}

func (f *fakeUserRepository) DeleteUser(id int64) error {
	f.deleteUserCalledWith = id
	return f.deleteUserErr
}

func TestUserServiceListUsers(t *testing.T) {
	repo := &fakeUserRepository{
		listUsersResult: []User{
			{ID: 1, Name: "Tom"},
			{ID: 2, Name: "Rose"},
		},
	}

	service := NewUserService(repo)

	users, err := service.ListUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}

func TestUserServiceGetUserByID(t *testing.T) {
	repo := &fakeUserRepository{
		getUserResult: User{ID: 10, Name: "Rose"},
	}

	service := NewUserService(repo)

	user, err := service.GetUserByID(10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if repo.getUserCalledWith != 10 {
		t.Fatalf("expected repo to be called with 10, got %d", repo.getUserCalledWith)
	}

	if user.Name != "Rose" {
		t.Fatalf("expected Rose, got %s", user.Name)
	}
}

func TestUserServiceDeleteUserNotFound(t *testing.T) {
	repo := &fakeUserRepository{
		deleteUserErr: ErrUserNotFound,
	}

	service := NewUserService(repo)

	err := service.DeleteUser(99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}
