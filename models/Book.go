package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Identifier string   `json:"identifier"`
	Authors    []Author `gorm:"many2many:book_authors;"`
	Title      string   `json:"title"`
	Subtitle   string   `json:"subtitle"`
}
