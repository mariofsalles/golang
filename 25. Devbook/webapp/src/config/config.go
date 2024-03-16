package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// ApiUrl represents the URL api calls are made to
	APIURL = ""
	// Port represents the port the application runs on
	Port = 0
	// HashKey is the key used to authenticate the cookie
	HashKey []byte
	// BlockKey is the key used to encrypt data in the cookie
	BlockKey []byte
)

// Load the enviornment variables
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("error loading PORT: %v", err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
