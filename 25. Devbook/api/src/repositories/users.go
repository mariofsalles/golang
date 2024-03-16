package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users is the repository for users
type Users struct {
	db *sql.DB
}

// NewUsersRepository creates a users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create an user that will be included on database
func (repository Users) Create(user models.User) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO users (username, nick, email, userpass) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var lastInsertedID uint64
	err = stmt.QueryRow(user.Username, user.Nick, user.Email, user.Password).Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

// Search for users by username or nick
func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, err := repository.db.Query(
		"SELECT id, username, nick, email, created_at FROM users WHERE username LIKE $1 OR nick LIKE $2", nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// SearchByID returns an user by its ID
func (repository Users) SearchByID(ID uint64) (models.User, error) {
	row, err := repository.db.Query("SELECT id, username, nick, email, created_at FROM users WHERE id = $1", ID)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// SearchByEmail returns an user by its email
func (repository Users) SearchByEmail(email string) (models.User, error) {
	row, err := repository.db.Query("SELECT id, userpass FROM users WHERE email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Update updates an user on database
func (repository Users) Update(userID uint64, user models.User) error {
	stmt, err := repository.db.Prepare("UPDATE users SET username = $1, nick = $2, email = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Username, user.Nick, user.Email, userID); err != nil {
		return err
	}
	return nil
}

// Delete removes an user from database
func (repository Users) Delete(userID uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID); err != nil {
		return err
	}

	return nil
}

// Follow allows an user to follow another
func (repository Users) Follow(userID, followerID uint64) error {
	stmt, err := repository.db.Prepare(
		"INSERT INTO followers (user_id, follower_id) VALUES ($1, $2) ON CONFLICT (user_id, follower_id) DO NOTHING")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow allows an user to unfollow another
func (repository Users) Unfollow(userID, followerID uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM followers WHERE user_id = $1 AND follower_id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Followers returns all followers from an user
func (repository Users) Followers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.username, u.nick, u.email, u.created_at
			FROM users u 
			INNER JOIN followers f 
			ON u.id = f.follower_id 
			WHERE f.user_id = $1`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var follower models.User
		if err = rows.Scan(
			&follower.ID,
			&follower.Username,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}
	return followers, nil
}

// Following returns all users that an user is following
func (repository Users) Following(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT u.id, u.username, u.nick, u.email, u.created_at
			FROM users u 
			INNER JOIN followers f 
			ON u.id = f.user_id 
			WHERE f.follower_id = $1`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var following []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		following = append(following, user)
	}
	return following, nil
}

// SearchPassword returns the password of an user
func (repository Users) SearchPassword(userID uint64) (string, error) {
	row, err := repository.db.Query("SELECT userpass FROM users WHERE id = $1", userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var user models.User
	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}
	return user.Password, nil
}

// UpdatePassword updates the password of an user
func (repository Users) UpdatePassword(userID uint64, userpass string) error {
	stmt, err := repository.db.Prepare("UPDATE users SET userpass = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(userpass, userID); err != nil {
		return err
	}
	return nil
}
