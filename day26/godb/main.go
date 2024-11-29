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
	ID       uint
	Name     string `gorm:"unique"`
	Age      int
	Birthday time.Time
}

func main() {
	// user:pwd@tcp(127.0.0.1:3306)/db_name
	URL := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PWD") + "@tcp(127.0.0.1:3306)/moshdb"
	db := ConnectDB(URL)

	var user User

	db.Model(&User{}).Find(&user)
	// db.Where("name Like ?", "%h%").Find(&user)
	// db.Where(&User{Name: "Sinzhu"}).Find(&user)
	// db.Find(&user)
	// db.Find(&user)

	// db.Where("name LIKE ?", "%hu%").First(&user)

	// fmt.Printf("%v\n", user)

	// user.Name = "Don"
	// user.Age = 34
	// db.Save(&User{ID: 7, Name: "Don"})
	db.Model(&User{}).Where("id = ?", 7).Update("name", "Nightwing")
	db.Model(&User{}).Where("id = ?", 13).Update("name", "Wonderwoman")
	db.Model(&User{}).Where("id = ?", 1).Update("name", "Batman")
	db.Model(&User{}).Where("id = ?", 15).Update("name", "Superman")

	db.Save(&User{ID: 1, Name: "BeastBoy", Age: 19})

	var userOne User
	db.First(&userOne)

	userOne.Name = "MrNobody"
	userOne.Age = 39

	fmt.Println(userOne.Name)
	// db.Delete(&userOne)
	db.Delete(&User{ID: 7})
}

func ConnectDB(URL string) gorm.DB {
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to db...")
	}
	return *db
}
