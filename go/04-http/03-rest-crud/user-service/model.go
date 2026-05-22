package main

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

var users = []User{
	{
		ID:   1,
		Name: "Tom",
	},
	{
		ID:   2,
		Name: "Rose",
	},
}
