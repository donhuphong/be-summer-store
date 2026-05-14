package main

import (
	"be-summer-store/internal/config"
	"be-summer-store/internal/controller"
	"be-summer-store/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.InitDB(cfg.DatabaseDSN)

	//database.DB.AutoMigrate(&model.Product{})

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/products", controller.GetProducts)
		api.GET("/presign",    controller.GetPresignURL)
		api.POST("/products",  controller.CreateProduct)
	}

	r.Run(":" + cfg.Port)
}
