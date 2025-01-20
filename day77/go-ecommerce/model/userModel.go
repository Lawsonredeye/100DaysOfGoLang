package model

import "time"

type Users struct {
	ID           uint
	Username     string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Role         string `json:"role"`
	Created_at   time.Time
	Updated_at   time.Time
}

type Product struct {
	ID           uint
	Name         string
	Description  string
	sku          string
	category_id  int
	color        string
	product_size int
	price        float64
	quantity     int
	created_at   time.Time
	updated_at   time.Time
}
