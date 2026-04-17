package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DbConnString is the connection string to postgres
	DbConnString = ""

	// Port is the port where the server will run
	Port = 0

	// SecretKey is the key to assign the token
	SecretKey []byte
)

// LoadingConfig loads the environment variables
func LoadingConfigs() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	DbConnString = fmt.Sprintf(
		`postgres://%s:%s@%s/%s?sslmode=disable`,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

}
