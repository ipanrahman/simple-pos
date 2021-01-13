package category

import (
	"github.com/gin-gonic/gin"
	"github.com/ipan97/simple-pos/internal/category/delivery/http"
	"github.com/ipan97/simple-pos/internal/category/repository"
	"gorm.io/gorm"
)

// InitRouter category
func InitRouter(v1 *gin.RouterGroup, db *gorm.DB) {
	categoryRepository := repository.NewCategoryRepository(db)
	categoryHandler := http.NewCategoryHTTPHandler(categoryRepository)
	v1.GET("/categories", categoryHandler.GetAllCategories)
	v1.GET("/categories/:id", categoryHandler.GetCategoryByID)
}
