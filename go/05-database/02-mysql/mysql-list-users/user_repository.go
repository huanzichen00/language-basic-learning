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

func clearUsersTable(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		return fmt.Errorf("clear users table failed: %w", err)
	}

	return nil
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

func updateUser(db *sql.DB, id int64, name string) error {
	result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return fmt.Errorf("update user failed: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected failed: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func deleteUser(db *sql.DB, id int64) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete user failed: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected failed: %w", err)
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
