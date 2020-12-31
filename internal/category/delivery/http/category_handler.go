package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ipan97/simple-pos/internal/core"
)

type CategoryHTTPHandler struct {
	core.BaseHandler
}

func (h *CategoryHTTPHandler) GetAllCategories(ctx *gin.Context) {
	h.OK(ctx, "Success get all categories", gin.H{})
}
