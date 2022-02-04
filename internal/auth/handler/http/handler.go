package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func (h *Handler) Login(ctx *gin.Context) {
	data := map[string]interface{}{
		"username": "Test",
		"password": "Test",
	}
	var result TestData
	j, err := json.Marshal(&data)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(j, &result); err != nil {
		log.Printf("Error : %v \n", err)
		return
	}
	ctx.JSON(http.StatusOK, &result)
}
