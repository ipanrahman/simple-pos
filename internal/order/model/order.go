package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID string `gorm:"column:order_id"`
}
