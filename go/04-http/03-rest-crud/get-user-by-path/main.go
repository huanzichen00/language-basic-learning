package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
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

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
}

func main() {
	http.HandleFunc("/users/", userByIDHandler)

	fmt.Println("server started at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
