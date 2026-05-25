package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	fmt.Println("path:", r.URL.Path)

	fmt.Fprintln(w, "hello, go server")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("server started at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}
