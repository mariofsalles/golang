package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

// Logger write information about the request on the console
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(": %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Authentication is a middleware to validate the token from user
func Authentication(nextFunc http.HandlerFunc) http.HandlerFunc {
	// http.HandlerFunc is a type that represents func (w http.ResponseWriter, r *http.Request)
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.TokenValidating(r); err != nil {
			responses.ErrJSON(w, http.StatusUnauthorized, err)
			return
		}
		// nextFunc(w, r) is the function that will be called after the middleware
		nextFunc(w, r)
	}
}
