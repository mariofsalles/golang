package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// Logout removes the user's cookie and redirects to the login page
func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
