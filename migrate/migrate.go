package main

import (
	"github.com/rovezuka/go-crud/initializers"
	"github.com/rovezuka/go-crud/models"
)

func init() {
	// Load environment variables and connect to the database
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Database migration
	initializers.DB.AutoMigrate(&models.Post{})
}
