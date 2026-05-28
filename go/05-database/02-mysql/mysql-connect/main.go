package main

import (
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

	id, err := createUser(db, "Tom")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	user, err := findUserByID(db, id)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("user id:", user.ID)
	fmt.Println("user name:", user.Name)

	userList, err := listUsers(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, user := range userList {
		fmt.Println("list user:", user.ID, user.Name)
	}
}
