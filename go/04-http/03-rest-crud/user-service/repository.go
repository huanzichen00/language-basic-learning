package main

type userRepository interface {
	List() []User
	GetByID(id int) (User, bool)
	Create(name string) User
	Update(id int, name string) (User, bool)
	Delete(id int) bool
}

type InMemoryUserRepository struct{}

func (r InMemoryUserRepository) List() []User {
	return users
}

func (r InMemoryUserRepository) GetByID(id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}

	return User{}, false
}

func (r InMemoryUserRepository) Create(name string) User {
	newUser := User{
		ID:   len(users) + 1,
		Name: name,
	}

	users = append(users, newUser)

	return newUser
}

func (r InMemoryUserRepository) Update(id int, name string) (User, bool) {
	for i := range users {
		if users[i].ID == id {
			users[i].Name = name
			return users[i], true
		}
	}

	return User{}, false
}

func (r InMemoryUserRepository) Delete(id int) bool {
	for i := range users {
		if users[i].ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}

	return false
}
