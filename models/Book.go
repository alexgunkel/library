package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Identifier string `json:"identifier"`
	Author string `json:"author"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
}
