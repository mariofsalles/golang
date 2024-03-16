package models

// AuthData is the struct that represents the data returned by the API
type AuthData struct {
	ID  string `json:"id"`
	Token string `json:"token"`
}