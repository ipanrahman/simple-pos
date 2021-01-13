package main

import (
	"fmt"
	"github.com/ipan97/simple-pos/internal/brand"
	"github.com/ipan97/simple-pos/internal/category"
	categoryModel "github.com/ipan97/simple-pos/internal/category/model"
	"github.com/ipan97/simple-pos/internal/customer"
	"github.com/ipan97/simple-pos/internal/product"
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
	db, errConnectDB := config.ConnectPostgres(conf)
	if errConnectDB != nil {
		log.Fatalf("Can't connect postgres cause : %v", errConnectDB)
	}
	errDBAutoMigrate := db.AutoMigrate(
		brand.Brand{},
		categoryModel.Category{},
		customer.Customer{},
		product.Product{},
	)
	if errDBAutoMigrate != nil {
		log.Fatalf("Can't migrate database : %v", errDBAutoMigrate)
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
	v1 := r.Group("/v1")
	{
		// Category Route
		category.InitRouter(v1, db)
	}
	if err := r.Run(fmt.Sprintf(":%d", conf.AppPort)); err != nil {
		fmt.Printf("Error running server cause %v : ", err)
	}
}
