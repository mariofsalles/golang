package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://golang:golang@172.17.0.2/devbook?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error on open the database connection: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Error on ping the database: %v", err)
	}

	fmt.Println("Connected to the database")

	lines, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		log.Fatalf("Error on query the database: %v", err)
	}
	defer lines.Close()

	fmt.Println(lines.Columns())
}
