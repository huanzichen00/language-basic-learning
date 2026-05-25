package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	idText := r.URL.Query().Get("id")
	if idText == "" {
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, "failed to encode users", http.StatusInternalServerError)
			return
		}

		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
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
	http.HandleFunc("/users", userHandler)

	fmt.Println("server started at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
