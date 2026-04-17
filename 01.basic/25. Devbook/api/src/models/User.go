package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user in the system
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"userpass,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will format and validate the user data
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}
	if err := user.formatter(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Username == "" {
		return errors.New("valid username is required")
	}
	if user.Nick == "" {
		return errors.New("valid nick is required")
	}
	if user.Email == "" {
		return errors.New("the email is not be blank")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("the email is not valid")
	}
	if stage == "createUser" && user.Password == "" {
		return errors.New("valid password is required")
	}
	return nil
}

func (user *User) formatter(stage string) error {
	user.Username = strings.TrimSpace(user.Username)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if stage == "createUser" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
	}
	return nil
}
