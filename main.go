package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipan97/simple-pos/internal/config"
)

var (
	conf *config.Config
)

func main() {
	conf = config.NewConfig()
	_, errDB := config.ConnectPostgres(conf)
	if errDB != nil {
		log.Fatalf("Can't connect postgres cause : %v", errDB)
	}
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
