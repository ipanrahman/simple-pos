package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string  `json:"product_name"`
	BrandID     string  `json:"brand_id"`
	CategoryID  string  `json:"category_id"`
	ModelYear   int     `json:"model_year"`
	Price       float32 `json:"price"`
}
