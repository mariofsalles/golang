package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger is a middleware to log all requests on terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(": %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authenticate is a middleware to check if the user has cookie
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.Read(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}
		next(w, r)
	}
}
