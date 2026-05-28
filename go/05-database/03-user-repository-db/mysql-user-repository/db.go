package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getEnvOrDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func openDB() (*sql.DB, error) {
	user := getEnvOrDefault("MYSQL_USER", "root")
	password := os.Getenv("MYSQL_PASSWORD")
	host := getEnvOrDefault("MYSQL_HOST", "127.0.0.1")
	port := getEnvOrDefault("MYSQL_PORT", "3306")
	dbName := getEnvOrDefault("MYSQL_DB", "go_basic_learning")

	if password == "" {
		return nil, fmt.Errorf("missing MYSQL_PASSWORD environment variable")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db failed: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping db failed: %w", err)
	}

	return db, nil
}

func createUsersTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id BIGINT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(100) NOT NULL
	)
	`
	_, err := db.Exec(query)
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
