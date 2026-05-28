package main

import (
	"database/sql"
	"errors"
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
func createAccount(db *sql.DB, name string, balance int64) (int64, error) {
	result, err := db.Exec("INSERT INTO accounts (name, balance) VALUES(?, ?)", name, balance)
	if err != nil {
		return 0, fmt.Errorf("create account failed: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id failed: %w", err)
	}

	return id, nil
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

func transferBalance(db *sql.DB, fromID int64, toID int64, amount int64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx failed: %w", err)
	}
	defer tx.Rollback()

	var fromBalance int64

	row := tx.QueryRow("SELECT balance FROM accounts WHERE id = ?", fromID)
	err = row.Scan(&fromBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("from account not found")
		}

		return fmt.Errorf("query from account balance failed: %w", err)
	}

	if fromBalance < amount {
		return fmt.Errorf("insufficient balance")
	}

	result, err := tx.Exec(
		"UPDATE accounts SET balance = balance - ? WHERE id = ?",
		amount,
		fromID,
	)
	if err != nil {
		return fmt.Errorf("decrease balace failed: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected for to account failed: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("from account not found")
	}

	result, err = tx.Exec(
		"UPDATE accounts SET balance = balance + ? WHERE id = ?",
		amount,
		toID,
	)

	affected, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("increase balance failed: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("to account cannot found")
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx failed: %w", err)
	}

	return nil
}

func printAccounts(title string, accounts []Account) {
	fmt.Println(title)
	for _, account := range accounts {
		fmt.Println("account:", account.ID, account.Name, account.Balance)
	}
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
	}

	tomID, err := createAccount(db, "Tom", 1000)
	if err != nil {
		fmt.Println("error:", err)
	}

	//RoseID, err := createAccount(db, "Rose", 500)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}

	accounts, err := listAccounts(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	printAccounts("before transfer:", accounts)

	err = transferBalance(db, tomID, 99999, 200)
	if err != nil {
		fmt.Println("error:", err)
	}

	accounts, err = listAccounts(db)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	printAccounts("after transfer:", accounts)
}
