package main

import (
	"fmt"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func loggingMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		fmt.Println("started", r.Method, r.URL.Path)

		next(recorder, r)

		fmt.Println(
			"finished",
			r.Method,
			r.URL.Path,
			"status",
			recorder.statusCode,
			"in",
			time.Since(start),
		)
	}
}
