package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ipan97/simple-pos/internal/category/repository"
	"github.com/ipan97/simple-pos/internal/core"
)

type CategoryHTTPHandler struct {
	core.BaseHandler
	categoryRepository repository.Repository
}

func NewCategoryHTTPHandler(categoryRepository repository.Repository) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{categoryRepository: categoryRepository}
}

func (h *CategoryHTTPHandler) GetAllCategories(ctx *gin.Context) {
	categories, err := h.categoryRepository.FindAll()
	if err != nil {
		h.NotFound(ctx, err.Error())
	}
	h.OK(ctx, "Success get all categories", categories)
}
