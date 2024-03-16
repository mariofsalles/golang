package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

// Login uses the email and password to authenticate the user
func LoginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	loggedUser, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"userpass": r.FormValue("userpass"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest,
			responses.ErrorAPI{
				Error: err.Error(),
			})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(loggedUser))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError,
			responses.ErrorAPI{
				Error: err.Error(),
			})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.StatusCodeErrorHandled(w, response)
		return
	}

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity,
			responses.ErrorAPI{
				Error: err.Error(),
			})
		return
	}

	if err = cookies.Save(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity,
			responses.ErrorAPI{
				Error: err.Error(),
			})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
