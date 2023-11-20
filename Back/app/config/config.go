package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// # Server settings:
	SERVER_URL=""
	SERVER_READ_TIMEOUT=60
	
	// # JWT settings:
	// # JWT_SECRET_KEY="secret"
	// # JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=15
	
	// # Database settings:
	DB_SERVER_URL=""
	DB_MAX_CONNECTIONS=100
	DB_MAX_IDLE_CONNECTIONS=10
	DB_MAX_LIFETIME_CONNECTIONS=2
)


func LoadEnvironmentVariables() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	SERVER_URL = os.Getenv("SERVER_URL")
	DB_SERVER_URL = os.Getenv("DB_SERVER_URL")
}