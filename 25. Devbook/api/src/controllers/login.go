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
	"io"
	"net/http"
	"strconv"
)

// Login make the login of an user
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ErrJSON(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
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
	userOnDB, err := repository.SearchByEmail(user.Email)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userOnDB.Password, user.Password); err != nil {
		responses.ErrJSON(w, http.StatusUnauthorized, errors.New("something is wrong with your access"))
		return
	}

	token, err := authentication.CreateToken(userOnDB.ID)
	if err != nil {
		responses.ErrJSON(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatUint(userOnDB.ID, 10)
	responses.JSON(w, http.StatusOK, models.AuthData{ID: userID, Token: token})
}
