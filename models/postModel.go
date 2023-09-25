package models

import "gorm.io/gorm"

// Модель для базы данных
type Post struct {
	gorm.Model
	Title string
	Body  string
}
