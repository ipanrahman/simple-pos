package http

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ipan97/simple-pos/internal/category/repository"
	"github.com/ipan97/simple-pos/internal/core"
	"gorm.io/gorm"
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
		h.InternalServerError(ctx, err)
		return
	}
	h.OK(ctx, "Success get all categories", categories)
}

func (h *CategoryHTTPHandler) GetCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, _ := strconv.Atoi(id)
	category, err := h.categoryRepository.FindOneByID(uint(categoryID))
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			h.NotFound(ctx, err.Error())
			return
		}
		h.InternalServerError(ctx, err)
		return
	}
	h.OK(ctx, fmt.Sprintf("Success get category by id : %s", id), category)
}
