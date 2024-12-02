package main

import (
	"log"
	"project-root/gptdb/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "lawsonredeye:lawson007@tcp(127.0.0.1)/gpt?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to db, %v", err)
	}

	db.AutoMigrate(&models.User{})

	// user := models.User{Name: "sean", Password: "NotDiddy", Email: "sean@mail.com"}
	// db.Create(&user)
	readUser(db, "lawson@mail.com")
	// updateUser(db, "sean@mail.com", "bigsean")
	// updateUser(db, "lawson@mail.com", "lawsonredeye")
	deleteUser(db, "lawson@mail.com")
}

func readUser(db *gorm.DB, email string) {
	var user models.User

	result := db.First(&user, "email = ?", email)

	if result.Error != nil {
		log.Println("User not found:", result.Error)
		return
	}

	log.Println("User found:", user)
}

func updateUser(db *gorm.DB, email, newName string) {
	var user models.User

	result := db.First(&user, "email = ?", email)

	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	user.Name = newName

	result = db.Save(&user)
	if result.Error != nil {
		log.Println("User was not updated, ", result.Error)
		return
	}
	log.Printf("user %v updated.", user.ID)
}

func deleteUser(db *gorm.DB, email string) {
	var user models.User

	result := db.First(&user, "email = ?", email)
	if result.Error != nil {
		log.Printf("%e", result.Error)
		return
	}

	db.Unscoped().Delete(&user)
	log.Println(email, "successfully deleted.")
}
