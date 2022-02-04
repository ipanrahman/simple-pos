package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
}

// OK
func (BaseHandler) OK(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, &APIResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// NotFound
func (BaseHandler) NotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, &APIResponse{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

// InternalServerError
func (BaseHandler) InternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, &APIResponse{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Error: &APIError{
			Message: err.Error(),
		},
	})
}
