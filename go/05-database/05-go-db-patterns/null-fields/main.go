package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64
	Name     string
	Nickname sql.NullString
	Age      sql.NullInt64
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		return fmt.Errorf("failed to drop table users: %w", err)
	}

	query := `
		CREATE TABLE IF NOT EXISTS users (
			id BIGINT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(100) NOT NULL,
			nickname VARCHAR(100) NULL,
			AGE BIGINT NULL
		)
	`

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("create users table failed: %w", err)
	}

	return nil
}

func clearUsersTable(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		return fmt.Errorf("clear users table failed: %w", err)
	}

	return nil
}

func insertUsers(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO users(name, nickname, age) VALUES(?, ?, ?)",
		"Tom",
		"Tommy",
		20,
	)
	if err != nil {
		return fmt.Errorf("insert Tom failed: %w", err)
	}

	_, err = db.Exec(
		"INSERT INTO users(name, nickname, age) VALUES(?, ?, ?)",
		"Rose",
		nil,
		nil,
	)
	if err != nil {
		return fmt.Errorf("insert Rose failed: %w", err)
	}

	return nil
}

func listUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, nickname, age FROM users ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("query users failed: %w", err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Nickname, &user.Age)
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

func printUsers(users []User) {
	for _, user := range users {
		fmt.Println("id:", user.ID)
		fmt.Println("name:", user.Name)

		if user.Nickname.Valid {
			fmt.Println("nickname:", user.Nickname.String)
		} else {
			fmt.Println("nickname NULL")
		}

		if user.Age.Valid {
			fmt.Println("age:", user.Age.Int64)
		} else {
			fmt.Println("age NULL")
		}

		fmt.Println("---")
	}
}

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

	err = insertUsers(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	users, err := listUsers(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	printUsers(users)
}
