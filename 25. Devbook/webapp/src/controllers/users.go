package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CreateUser fetch the data from the form and send to the API
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"username": r.FormValue("username"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"userpass": r.FormValue("userpass"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	userURL := fmt.Sprintf("%s/users", config.APIURL)
	response, err := http.Post(userURL, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}
	responses.JSON(w, response.StatusCode, nil)
}

// GetUsersPage fetch the users from the API
func GetUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("username"))
	url := fmt.Sprintf("%s/users?username=%s", config.APIURL, nameOrNick)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}

	var users []models.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

// UnfollowUser send a request to the API to unfollow a user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// FollowUser send a request to the API to follow a user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// CreateUser fetch the data from the form and send to the API
func UpdateCurrentUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"username": r.FormValue("username"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPut, url, bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}
	responses.JSON(w, response.StatusCode, nil)
}

// UpdatePassword send a request to the API to update the user password
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newPassword, err := json.Marshal(map[string]string{
		"current_password": r.FormValue("current_password"),
		"new_password":     r.FormValue("new_password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/update-userpass", config.APIURL, userID)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodPost, url, bytes.NewBuffer(newPassword))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}
	responses.JSON(w, response.StatusCode, nil)
}

// DeleteUser send a request to the API to delete the user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, err := requests.MakeAuthenticatedRequest(r, http.MethodDelete, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}
	cookies.Delete(w)

	responses.JSON(w, response.StatusCode, nil)
}