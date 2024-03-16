package controllers

import (
	"api/src/authentication"
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnprocessableEntity, err)
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("createUser"); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user.ID, err = repository.Create(user)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers returns all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("username"))

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUser returns an user
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	user, err := repository.SearchByID(userId)

	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDInToken {
		responses.ErrJSON(w, http.StatusForbidden, fmt.Errorf("you can only update your own user"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("updateUser"); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	if err = repository.Update(userID, user); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDInToken {
		responses.ErrJSON(w, http.StatusForbidden, fmt.Errorf("you can only delete your own user"))
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	if err = repository.Delete(userID); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser follows an user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	if followerID == userID {
		responses.ErrJSON(w, http.StatusForbidden, errors.New("you cannot follow yourself"))
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Follow(userID, followerID); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnFollowUser unfollows an user
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followID, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	if followID == userID {
		responses.ErrJSON(w, http.StatusForbidden, errors.New("wait, you never followed yourself"))
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Unfollow(userID, followID); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Followers returns all followers of an user
func Followers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	followers, err := repository.Followers(userID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// Following returns all users that an user follows
func Following(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	following, err := repository.Following(userID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

// UpdatePassword updates the password of an user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	userIDInToken, err := authentication.ExtractUserID(r)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDInToken {
		responses.ErrJSON(w, http.StatusForbidden, fmt.Errorf("you can only update your own user"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userpass models.Password
	if err = json.Unmarshal(requestBody, &userpass); err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConectOnDB()
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	passwordInDB, err := repository.SearchPassword(userID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(passwordInDB, userpass.CurrentPassword); err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, errors.New("the current password is incorrect"))
		return
	}

	passwordHash, err := security.Hash(userpass.NewPassword)
	if err != nil {
		responses.ErrJSON(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.UpdatePassword(userID, string(passwordHash)); err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
