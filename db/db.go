package db

import (
	"gin-gorm-curd-rest-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(){
	_db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db = _db
  if err != nil {
    panic("failed to connect database")
  }
	db.AutoMigrate(&models.Product{})

}

func GetDB() *gorm.DB {
	return db
}