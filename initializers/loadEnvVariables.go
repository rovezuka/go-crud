package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Загрузка и получение переменных окружения
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
