package server

import (
	"crud/db"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint64 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Prepare de request body
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Error reading request body\"}"))
		return
	}

	// entity used in request body
	var user user

	// Write values from struct into json format used by request body
	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"message\":\"Error unmarshaling request\"}"))
		return
	}

	// connect to the database
	db, err := db.ConectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error connecting to database\"}"))
		return
	}
	defer db.Close()

	// prepare the statement
	stmt, err := db.Prepare("INSERT INTO usuarios(nome, email) VALUES($1, $2)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error preparing the statement\"}"))
		return
	}
	defer stmt.Close()

	// execute the statement
	insert, err := stmt.Exec(user.Nome, user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error executing the statement\"}"))
		return
	}

	// return if the user was created successfully
	rowsAffected, err := insert.RowsAffected()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error getting the number of rows affected\"}"))
		return
	}

	if rowsAffected == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"No rows were affected by the insert operation\"}"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{\"message\":\"User created successfully\"}"))
}

// GetUsers get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Open connecting from database
	db, err := db.ConectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error connecting to database\"}"))
	}
	defer db.Close()

	// prepare the statement
	lines, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error to get users\"}"))
		return
	}
	defer lines.Close()

	var users []user

	for lines.Next() {
		var user user
		if err := lines.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"message\":\"Error to scan user\"}"))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error to encode users from json\"}"))
	}
}

// GetUser get a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Because ID is a string, we need to convert it to int
	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error converting ID to int\"}"))
		return
	}

	// Open connecting from database
	db, err := db.ConectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error connecting to database\"}"))
		return
	}
	defer db.Close()

	// Make the query by ID
	line, err := db.Query("SELECT * FROM usuarios WHERE id = $1", ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error to get user\"}"))
		return
	}
	defer line.Close()

	var user user

	// Scan the user from the database by ID from the query
	if line.Next() {
		// If the user exists, we will record on user variable
		if err := line.Scan(&user.ID, &user.Nome, &user.Email); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"message\":\"Error to scan user\"}"))
			return
		}
	}

	// If the user does not exist, we will return a 404 status
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("{\"message\":\"User not found\"}"))
		return
	}

	// If the user exists, we will return it
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error to encode user from json\"}"))
	}
}

// UpdateUser update a user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error converting ID to int\"}"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error reading request body\"}"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error unmarshaling request\"}"))
		return
	}

	db, err := db.ConectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error connecting to database\"}"))
		return
	}
	defer db.Close()

	stm, err := db.Prepare("UPDATE usuarios SET nome = $1, email = $2 WHERE id = $3")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error preparing the statement\"}"))
		return
	}
	defer stm.Close()

	if _, err := stm.Exec(user.Nome, user.Email, ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error executing the statement\"}"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser delete a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error converting ID to int\"}"))
	}

	db, err := db.ConectionDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error connecting to database\"}"))
		return
	}
	defer db.Close()

	stm, err := db.Prepare("DELETE FROM usuarios WHERE id = $1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error preparing the statement\"}"))
		return
	}
	defer stm.Close()

	if _, err := stm.Exec(ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\":\"Error executing the statement\"}"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}