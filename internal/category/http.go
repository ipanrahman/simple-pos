package category

import (
	"github.com/gin-gonic/gin"
	"github.com/ipan97/simple-pos/internal/core"
)

type HTTPHandler struct {
	core.BaseHandler
}

func (h HTTPHandler) GetAllCategories(ctx *gin.Context) {
	h.OK(ctx, "Success get all categories", gin.H{})
}
