package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", loggingMiddleWare(usersHandler))
	http.HandleFunc("/users/", loggingMiddleWare(userByIDHandler))

	fmt.Println("server started at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error", err)
	}
}
