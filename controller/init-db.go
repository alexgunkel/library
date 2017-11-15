package controller

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/alexgunkel/library/models"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:very-secure@/database?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Author{})
}
