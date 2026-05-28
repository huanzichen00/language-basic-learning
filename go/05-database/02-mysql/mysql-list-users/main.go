package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func main() {
	db, err := openDB()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer db.Close()

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

	_, err = createUser(db, "Tom")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_, err = createUser(db, "Rose")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	_, err = createUser(db, "Jack")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("before update/delete:")

	userList, err := listUsers(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, user := range userList {
		fmt.Println("list user:", user.ID, user.Name)
	}

	err = updateUser(db, 2, "RoseUpdated")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("update target not found")
			return
		}

		fmt.Println("error:", err)
		return
	}

	err = deleteUser(db, 1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("delete target not found")
			return
		}

		fmt.Println("error:", err)
		return
	}

	fmt.Println("after update/delete:")

	userList, err = listUsers(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, user := range userList {
		fmt.Println("list user:", user.ID, user.Name)
	}
}
