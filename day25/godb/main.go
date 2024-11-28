package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

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

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&Product{Code: "r-643", Price: 291})

	// var product Product
	// result := db.First(&product, 1)
	// db.First(&product, "code = ?", "D42")

	// // update multiple field

	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "f42"})

	// db.Delete(&product, 1)

	// fmt.Println(result.RowsAffected)
}
