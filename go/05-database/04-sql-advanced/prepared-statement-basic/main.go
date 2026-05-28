package main

import (
	"database/sql"
	"fmt"
)

type Account struct {
	ID      int64
	Name    string
	Balance int64
}

func createAccountsTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS accounts (
            id BIGINT PRIMARY KEY AUTO_INCREMENT,
            name VARCHAR(100) NOT NULL,
            balance BIGINT NOT NULL
        )
    `

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("create accounts table failed: %w", err)
	}

	return nil
}

func clearAccountsTable(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM accounts")
	if err != nil {
		return fmt.Errorf("clear accounts table failed: %w", err)
	}

	return nil
}

func listAccounts(db *sql.DB) ([]Account, error) {
	rows, err := db.Query("SELECT id, name, balance FROM accounts ORDER BY id")
	if err != nil {
		return nil, fmt.Errorf("query accounts failed: %w", err)
	}
	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var account Account

		err := rows.Scan(&account.ID, &account.Name, &account.Balance)
		if err != nil {
			return nil, fmt.Errorf("scan account failed: %w", err)
		}

		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate accounts failed: %w", err)
	}

	return accounts, nil
}

func insertAccounts(db *sql.DB, accounts []Account) error {
	stmt, err := db.Prepare("INSERT INTO accounts(name, balance) VALUES(?, ?)")
	if err != nil {
		return fmt.Errorf("prepare insert account failed: %w", err)
	}
	defer stmt.Close()

	for _, account := range accounts {
		_, err := stmt.Exec(account.Name, account.Balance)
		if err != nil {
			return fmt.Errorf("exec insert account failed: %w", err)
		}
	}

	return nil
}

func main() {
	db, err := openDB()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer db.Close()

	err = createAccountsTable(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = clearAccountsTable(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	accountsToInsert := []Account{
		{
			Name:    "Tom",
			Balance: 1000,
		},
		{
			Name:    "Rose",
			Balance: 500,
		},
		{
			Name:    "Jack",
			Balance: 800,
		},
	}

	err = insertAccounts(db, accountsToInsert)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	accounts, err := listAccounts(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, account := range accounts {
		fmt.Println("account:", account.ID, account.Name, account.Balance)
	}
}
