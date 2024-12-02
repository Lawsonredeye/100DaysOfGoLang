package main

import (
	"project-root/db"
	"project-root/models"
	"testing"

	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := db.ConnectTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}
	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)

	user := models.User{Name: "John Doe", Email: "john@example.com"}
	db.Create(&user)

	var fetchedUser models.User
	db.First(&fetchedUser, "email = ?", "john@example.com")

	if fetchedUser.Name != "John Doe" {
		t.Errorf("Expected Name to be John Doe, got %v", fetchedUser.Name)
	}
}
