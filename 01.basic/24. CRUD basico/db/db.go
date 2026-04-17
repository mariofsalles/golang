package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// ConectarBD abre a conexão com o banco de dados
func ConectionDB() (*sql.DB, error) {
	connStr := "postgres://golang:golang@172.17.0.2/devbook?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
