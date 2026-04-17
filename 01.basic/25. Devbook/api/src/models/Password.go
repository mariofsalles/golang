package models

// Password represents the user password
type Password struct {
	NewPassword     string `json:"new_password,omitempty"`
	CurrentPassword string `json:"current_password,omitempty"`
}
