package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TotalValue float32     `json:"total_value"`
	OrderItems []OrderItem `json:"order_items"`
	UserId     int         `json:"user_id"`
	IsPayed    bool        `json:"payed"`
}
