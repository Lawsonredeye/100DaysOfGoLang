package models

import (
	"gorm.io/gorm"
)

// type User struct {
// 	ID        uint `gorm:"primarykey"`
// 	Name      string
// 	Email     string `gorm:"unique"`
// 	Password  string
// 	CreatedAt time.Time `gorm:"autoCreateTime"`
// 	UpdateAt  time.Time `gorm:"autoCreateTime"`
// }

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	// CreatedAt time.Time `gorm:"autoCreateTime"`
	// UpdatedAt time.Time `gorm:"autoCreateTime"`
}
