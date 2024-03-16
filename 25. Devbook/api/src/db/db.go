package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
)

// ConectionDB returns a connection to the database
func ConectOnDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DbConnString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
