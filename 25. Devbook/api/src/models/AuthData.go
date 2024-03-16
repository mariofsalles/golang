package models

// AuthData is the struct that represents the data returned from the login
type AuthData struct {
	ID 	string `json:"id"`
	Token string `json:"token"`
}