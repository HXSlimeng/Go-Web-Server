package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	database,err := gorm.Open(sqlite.Open("go-admin-db.db"),&gorm.Config{})
	if err !=nil {
		log.Fatal("err")
	}
	DB = database
}