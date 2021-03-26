package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username"  json:"username"`
	Email    string `gorm:"column:email"     json:"email"`
	Password string `gorm:"column:password"  json:"password"`
	IsActive bool   `gorm:"column:is_active" json:"is_active"`
}
