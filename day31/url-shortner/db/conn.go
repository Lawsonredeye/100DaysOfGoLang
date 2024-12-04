package db

import (
	"log"
	"os"
	"url-shortener/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var URL string = os.Getenv("db_user") + ":" + os.Getenv("db_pwd") + "@tcp(127.0.0.1)/shortlinkdb"

// Connects to the MySQL DB and performs auto migrations for you.
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to db, %v", err)
		return db, err
	}

	db.AutoMigrate(&model.ShortLink{})
	return db, nil
}
