package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	RoleName    string `gorm:"column:role_name;type:varchar(25)"`
	Description string `gorm:"column:description;type:varchar(50)"`
}
