package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type TestData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(ctx *gin.LoggerConfig) {
	data := map[string]interface{}{
		"username": "Test",
		"password": "Test",
	}
	var result TestData
	j, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(j, &result)

}
