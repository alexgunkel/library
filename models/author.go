package models

import (
	"github.com/jinzhu/gorm"
)

// Generic author type
type Author struct {
	gorm.Model
	FirstName string `json:"first_name" binding="required"`
	LastName  string `json:"last_name" binding="required"`
	Books     []Book `gorm:"many2many:book_authors;"`
}

func (a Author) String() string {
	return a.FirstName + " " + a.LastName
}
