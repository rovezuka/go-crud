package main

import (
	"github.com/rovezuka/go-crud/initializers"
	"github.com/rovezuka/go-crud/models"
)

func init() {
	// Загружаем переменные окружения и подключаемся к базе данных
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Миграция базы данных
	initializers.DB.AutoMigrate(&models.Post{})
}
