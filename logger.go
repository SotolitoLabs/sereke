package main

import (
	"log"
	"net/http"
	"time"
)

// Logger is a decorator that serves the request and creates an entry
// to the log file.
// TODO make it write to the log file :P
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s\n%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
			r.Header,
		)
	})
}
