package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = ConnectDB()

// ConnectDB connects to a MySQL database.
func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("HOST")
	passwd := os.Getenv("HOST_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dns := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v?parseTime=True", user, passwd, db_name)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic("Database connection was unsuccessful!\nCheck if MySQL is running.")
	}

	return db
}
