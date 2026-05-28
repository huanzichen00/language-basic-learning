package main

import "context"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) ListUsers(ctx context.Context) ([]User, error) {
	return s.repo.ListUsers(ctx)
}

func (s UserService) GetUserByID(ctx context.Context, id int64) (User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s UserService) CreateUser(ctx context.Context, name string) (User, error) {
	return s.repo.CreateUser(ctx, name)
}

func (s UserService) UpdateUser(ctx context.Context, id int64, name string) (User, error) {
	return s.repo.UpdateUser(ctx, id, name)
}

func (s UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}
