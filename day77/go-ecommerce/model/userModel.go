package model

import "time"

type Users struct {
	ID           uint
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	Role         string    `json:"role"`
	Created_at   time.Time `gorm:"autoCreateTime"`
	Updated_at   time.Time `gorm:"autoCreateTime"`
}

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	Category string `json:"category"`
	CategoryID  int       `json:"category_id"`
	Color       string    `json:"color"`
	ProductSize int       `json:"product_size"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Created_at  time.Time `gorm:"autoCreateTime"`
	Updated_at  time.Time `gorm:"autoCreateTime"`
}

type Categories struct {
	ID uint `json:"int"`
	Name string `json:"string"`
}
