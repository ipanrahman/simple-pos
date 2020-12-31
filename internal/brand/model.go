package brand

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	BrandID   string `json:"brand_id"`
	BrandName string `json:"brand_name"`
}
