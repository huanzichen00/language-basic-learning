package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var userService = UserService{
	repo: InMemoryUserRepository{},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		userList := userService.ListUsers()

		err := json.NewEncoder(w).Encode(userList)
		if err != nil {
			http.Error(w, "failed to encode users", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		var req CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid json body", http.StatusBadRequest)
			return
		}

		if req.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		newUser := userService.CreateUser(req.Name)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(newUser)
		if err != nil {
			http.Error(w, "failed to encode user", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "users" || parts[2] == "" {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		user, err := userService.GetUserByID(id)
		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "failed to encode user", http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		var req UpdateUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid json body", http.StatusBadRequest)
			return
		}

		if req.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		updateUser, err := userService.UpdateUser(id, req.Name)
		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(updateUser)
		if err != nil {
			http.Error(w, "failed to encode user", http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		err := userService.DeleteUser(id)
		if err != nil {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
