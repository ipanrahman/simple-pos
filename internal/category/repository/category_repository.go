package repository

import (
	"github.com/ipan97/simple-pos/internal/category/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Category, error)
	FindOneByID(id uint) (model.Category, error)
	FindAllByNameIgnoreCase(name string) ([]model.Category, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Model(&model.Category{}).Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) FindOneByID(id uint) (model.Category, error) {
	var category model.Category
	err := r.db.Model(&model.Category{}).Take(&category, "id=?", id).Error
	return category, err
}

func (r *CategoryRepository) FindAllByNameIgnoreCase(name string) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Model(&model.Category{}).Find(&categories, "name ILIKE ?", "%"+name+"%").Error
	return categories, err
}
