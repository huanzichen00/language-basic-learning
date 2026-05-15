package user

import "errors"

type UserRepository interface {
	FindByID(id int) (User, error)
}

type MemoryUserRepository struct {
	users map[int]User
}

func NewMemoryUserRepository(users map[int]User) MemoryUserRepository {
	return MemoryUserRepository{
		users: users,
	}
}

func (r MemoryUserRepository) FindByID(id int) (User, error) {
	user, ok := r.users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}
