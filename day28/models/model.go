package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// This package connects to a database for our program and initializes it for usage

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Bio      string
	Age      uint8
}

// Conn(DB_URL) connects to the database based on the mysql driver URI
// It returns a pointer to the db that can be used for running db operations.
func Conn(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	CheckDbError(err)
	db.AutoMigrate(&User{})
	return db
}

// Checks for the errors that may occur in during connection to our database.
func CheckDbError(e error) {
	if e != nil {
		log.Fatalf("%q", e)
	}
}
