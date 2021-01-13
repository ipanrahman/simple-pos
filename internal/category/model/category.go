package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"column:name"        json:"name"`
	Description string `gorm:"column:description" json:"description"`
}
