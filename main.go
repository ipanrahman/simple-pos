package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Masterminds/sprig"
	"github.com/gin-gonic/gin"
	"github.com/ipan97/simple-pos/internal/brand"
	"github.com/ipan97/simple-pos/internal/category"
	categoryModel "github.com/ipan97/simple-pos/internal/category/model"
	"github.com/ipan97/simple-pos/internal/config"
	"github.com/ipan97/simple-pos/internal/customer"
	"github.com/ipan97/simple-pos/internal/product/model"
	roleModel "github.com/ipan97/simple-pos/internal/role/model"
	userModel "github.com/ipan97/simple-pos/internal/user/model"
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
		model.Product{},
		userModel.User{},
		roleModel.Role{},
	)
	if errDBAutoMigrate != nil {
		log.Fatalf("Can't migrate database : %v", errDBAutoMigrate)
	}
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.Delims("{{", "}}")
	r.SetFuncMap(sprig.FuncMap())
	r.LoadHTMLGlob("web/*.html")

	// Main page
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "Simple Point Of Sales",
		})
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// Category Route
			category.InitRouter(v1, db)
		}
	}
	if err := r.Run(fmt.Sprintf(":%d", conf.AppPort)); err != nil {
		fmt.Printf("Error running server cause %v : ", err)
	}
}
