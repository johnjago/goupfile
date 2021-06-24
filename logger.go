package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%-5s\t%-10s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
