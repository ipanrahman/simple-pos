package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-pos/internal/config"
)

var (
	conf *config.Config
)

func init() {
	conf = config.NewConfig()
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Simple Point Of Sales")
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	if err := r.Run(fmt.Sprintf(":%d", conf.AppPort)); err != nil {
		fmt.Printf("Error running server cause %v : ", err)
	}
}
