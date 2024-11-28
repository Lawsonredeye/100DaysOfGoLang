package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	USERNAME := os.Getenv("user_name")
	PASSWD := os.Getenv("passwrd")
	connector := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/moshdb", USERNAME, PASSWD)
	db, err := gorm.Open(mysql.Open(connector), &gorm.Config{})

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("database failed to connect...")
	}
	log.Println("Database connection was successful...")

	// Creating Records
	type User struct {
		ID       uint
		Name     string `gorm:"unique"`
		Age      int
		Birthday time.Time
	}
	db.AutoMigrate(&User{})

	// user := []User{
	// 	{Name: "Lawson", Age: 43, Birthday: time.Now()},
	// 	{Name: "Kelvin", Age: 23, Birthday: time.Now()},
	// 	{Name: "Anthony", Age: 46, Birthday: time.Now()},
	// }

	// for _, val := range user {
	// 	result := db.Create(&val)

	// 	fmt.Println(val.ID)
	// 	fmt.Println("found errors:", result.Error)
	// }

	// Batch insert into db
	// db.Model(&User{}).Create(map[string]interface{}{
	// 	"Name": "Sinzhu Money", "Age": 32,
	// })

	// db.Model(&User{}).Create([]map[string]interface{}{
	// 	{"Name": "Jinzhu Smith", "Age": 21},
	// 	{"Name": "Kinzhu Kelly", "Age": 31},
	// })

	// Querying
	var user = User{ID: 12}
	result := db.Find(&user)
	fmt.Println(result.Error)
	if result.Error == nil {
		fmt.Println(user.Name)
	} else {
		log.Println("No user with such id")
	}
}
