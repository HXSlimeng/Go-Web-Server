package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	database,err := gorm.Open(sqlite.Open("D:/prtTSc/go-admin/config/db.sql"),&gorm.Config{})
	if err !=nil {
		log.Fatal("err")
	}
	DB = database
}