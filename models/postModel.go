package models

import "gorm.io/gorm"

// Model for the database
type Post struct {
	gorm.Model
	Title string
	Body  string
}
