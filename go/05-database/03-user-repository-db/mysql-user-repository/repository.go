package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	CreateUser(ctx context.Context, name string) (User, error)
	UpdateUser(ctx context.Context, id int64, name string) (User, error)
	DeleteUser(ctx context.Context, id int64) error
}

type MySQLUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return MySQLUserRepository{
		db: db,
	}
}

func (r MySQLUserRepository) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM users ORDER BY id")
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

func (r MySQLUserRepository) GetUserByID(ctx context.Context, id int64) (User, error) {
	var user User

	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ?", id)

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}

		return User{}, fmt.Errorf("scan user failed: %w", err)
	}

	return user, nil
}

func (r MySQLUserRepository) CreateUser(ctx context.Context, name string) (User, error) {
	result, err := r.db.ExecContext(ctx, "INSERT INTO users(name) VALUES(?)", name)
	if err != nil {
		return User{}, fmt.Errorf("insert user failed: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, fmt.Errorf("get last insert id failed: %w", err)
	}

	return User{
		ID:   id,
		Name: name,
	}, nil
}

func (r MySQLUserRepository) DeleteUser(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete user failed: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected failed: %w", err)
	}

	if affected == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (r MySQLUserRepository) UpdateUser(ctx context.Context, id int64, name string) (User, error) {
	result, err := r.db.ExecContext(ctx, "UPDATE users SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return User{}, fmt.Errorf("update user failed: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return User{}, fmt.Errorf("get rows affected failed: %w", err)
	}

	if affected == 0 {
		return User{}, ErrUserNotFound
	}

	return User{
		ID:   id,
		Name: name,
	}, nil
}
