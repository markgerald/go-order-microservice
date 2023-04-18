package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Amount  int     `json:"amount"`
	OrderID uint    `json:"order_id"`
}
