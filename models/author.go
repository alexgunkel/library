package models

import "github.com/jinzhu/gorm"

// Generic author type
type Author struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
