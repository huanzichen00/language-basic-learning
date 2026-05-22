package main

import (
	"fmt"
	"net/http"
	"time"
)

func loggingMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		fmt.Println("started", r.Method, r.URL.Path)

		next(w, r)

		fmt.Println("finished", r.Method, r.URL.Path, "in", time.Since(start))
	}
}
