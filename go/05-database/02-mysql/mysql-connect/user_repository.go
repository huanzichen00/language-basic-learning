package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID   int64
	Name string
}

func createUser(db *sql.DB, name string) (int64, error) {
	result, err := db.Exec("INSERT INTO users(name) VALUES(?)", name)
	if err != nil {
		return 0, fmt.Errorf("insert user failed: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id failed: %w", err)
	}

	return id, nil
}

func findUserByID(db *sql.DB, id int64) (User, error) {
	var user User

	row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, fmt.Errorf("user not found: %w", err)
		}

		return User{}, fmt.Errorf("scan user failed: %w", err)
	}

	return user, nil
}

func listUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name FROM users ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("query users failed: %w", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("scan user failed: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate users failed: %w", err)
	}

	return users, nil
}
