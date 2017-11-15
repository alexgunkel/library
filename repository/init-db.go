package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/alexgunkel/library/models"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Database struct {
		Username string
		Password string
		Database string
	}
}

var db *gorm.DB

func init() {
	db = getDatabase()
	db.LogMode(true)
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Author{})
}

func getDatabase() *gorm.DB {
	var err error
	var args string
	var cfg Config

	err = gcfg.ReadFileInto(&cfg, "config.cfg")
	if nil != err {
		panic(err)
	}

	args = cfg.Database.Username + ":" + cfg.Database.Password + "@/" + cfg.Database.Database + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open("mysql", args)

	if err != nil {
		panic(err)
	}

	return db
}
