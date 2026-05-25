package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			ID:   1,
			Name: "Tom",
		},
		{
			ID:   2,
			Name: "Rose",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "failed to encode users", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)

	fmt.Println("server started at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
