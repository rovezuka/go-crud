package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Loading and retrieving environment variables
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
