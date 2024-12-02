package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect to in-memory db for testing
func ConnectTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}
