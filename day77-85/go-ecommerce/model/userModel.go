package model

import "time"

type Users struct {
	ID           uint
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	Role         string    `json:"role"`
	Created_at   time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at   time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

type Product struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	Category    string    `json:"category"`
	CategoryID  int       `json:"category_id"`
	Color       string    `json:"color"`
	ProductSize int       `json:"product_size"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Created_at  time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at  time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

type Categories struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Abbrevation string `json:"abbrevation"`
}

type Orders struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	Created_at  time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at  time.Time `gorm:"autoCreateTime" json:"updated_at"`
}

type OrderItem struct {
	ID         uint      `json:"id"`
	OrderID    int       `json:"order_id"`
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Price      float64   `json:"price"`
	Created_at time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at time.Time `gorm:"autoCreateTime" json:"updated_at"`
}
