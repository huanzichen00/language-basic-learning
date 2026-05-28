package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	db, err := openDB()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = createUsersTable(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = clearUsersTable(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	repo := NewUserRepository(db)
	service := NewUserService(repo)

	tom, err := service.CreateUser(ctx, "Tom")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	rose, err := service.CreateUser(ctx, "Rose")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	users, err := service.ListUsers(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("before update/delete:")
	for _, user := range users {
		fmt.Println("user:", user.ID, user.Name)
	}

	updatedUser, err := service.UpdateUser(ctx, rose.ID, "RoseUpdated")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("updated:", updatedUser.ID, updatedUser.Name)

	err = service.DeleteUser(ctx, tom.ID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	users, err = service.ListUsers(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("after update/delete:")
	for _, user := range users {
		fmt.Println("user:", user.ID, user.Name)
	}

	_, err = service.GetUserByID(ctx, tom.ID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			fmt.Println("deleted user not found")
			return
		}

		fmt.Println("error:", err)
		return
	}
}
