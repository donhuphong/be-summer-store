package main

import (
	"be-summer-store/internal/config"
	"be-summer-store/internal/database"
	"be-summer-store/internal/router"
)

func main() {
	config.LoadConfig()

	database.InitDB()
	database.InitR2()

	//database.DB.AutoMigrate(&model.Product{})

	r := router.SetupRouter()

	r.Run(":" + config.AppConfig.Port)
}
