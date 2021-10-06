package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Login(ctx *gin.LoggerConfig) {
	data := map[string]interface{}{
		"username": "Test",
		"password": "Test",
	}
	json.Unmarshal(data, nil)
}
