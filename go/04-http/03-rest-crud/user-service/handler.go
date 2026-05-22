package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(users)
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

		newUser := User{
			ID:   len(users) + 1,
			Name: req.Name,
		}

		users = append(users, newUser)

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
		for _, user := range users {
			if user.ID == id {
				w.Header().Set("Content-Type", "application/json")

				err := json.NewEncoder(w).Encode(user)
				if err != nil {
					http.Error(w, "failed to encode user", http.StatusInternalServerError)
					return
				}

				return
			}
		}

		http.Error(w, "user not found", http.StatusNotFound)

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

		for i := range users {
			if users[i].ID == id {
				users[i].Name = req.Name

				w.Header().Set("Content-Type", "application/json")

				err := json.NewEncoder(w).Encode(users[i])
				if err != nil {
					http.Error(w, "failed to encode user", http.StatusInternalServerError)
					return
				}

				return
			}
		}

		http.Error(w, "user not found", http.StatusNotFound)

	case http.MethodDelete:
		for i := range users {
			if users[i].ID == id {
				users = append(users[:i], users[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		http.Error(w, "user not found", http.StatusNotFound)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
