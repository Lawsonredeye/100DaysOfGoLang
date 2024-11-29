package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `gorm:"name"`
	Birthday time.Time
	Age      int `gorm:"age"`
}

func main() {
	// user:pwd@tcp(127.0.0.1:3306)/db_name
	URL := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PWD") + "@tcp(127.0.0.1:3306)/moshdb"
	db := ConnectDB(URL)

	log.Println("DB connected successfully..")

	var user = User{Name: "lawson"}
	db.Model(&User{}).First(user)

	log.Println(user.Name)

	result := map[string]interface{}{}
	db.Table("users").Take(&result)

	fmt.Printf("%v\n", result)
}

func ConnectDB(URL string) gorm.DB {
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to db...")
	}
	return *db
}
