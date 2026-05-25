package main

import "testing"

func TestUserServiceListUsers(t *testing.T) {
	resetUsers()

	service := UserService{
		repo: InMemoryUserRepository{},
	}

	got := service.ListUsers()

	if len(got) != 2 {
		t.Fatalf("expected 2 users, got %d", len(got))
	}
}

func TestUserServiceGetUserID(t *testing.T) {
	resetUsers()

	service := UserService{
		repo: InMemoryUserRepository{},
	}

	user, err := service.GetUserByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "Tom" {
		t.Fatalf("expected Tom, got %s", user.Name)
	}
}

func TestUserServiceGetUserByIDNotFound(t *testing.T) {
	resetUsers()

	service := UserService{
		repo: InMemoryUserRepository{},
	}

	_, err := service.GetUserByID(99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUserServiceCreateUser(t *testing.T) {
	resetUsers()

	service := UserService{
		repo: InMemoryUserRepository{},
	}

	user := service.CreateUser("Jack")

	if user.ID != 3 {
		t.Fatalf("expected ID 3, got %d", user.ID)
	}

	if user.Name != "Jack" {
		t.Fatalf("expected Jack, got %s", user.Name)
	}
}

func TestUserServiceUpdateUser(t *testing.T) {
	resetUsers()

	service := UserService{
		repo: InMemoryUserRepository{},
	}

	user, err := service.UpdateUser(2, "RoseUpdated")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.Name != "RoseUpdated" {
		t.Fatalf("expected RoseUpdated, got %s", user.Name)
	}
}

func TestUserServiceUpdateUserNotFound(t *testing.T) {
	resetUsers()

	service := UserService{
		repo: InMemoryUserRepository{},
	}

	_, err := service.UpdateUser(99, "Nobody")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestUserServiceDeleteUser(t *testing.T) {
	resetUsers()

	service := UserService{
		repo: InMemoryUserRepository{},
	}

	err := service.DeleteUser(2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
}
