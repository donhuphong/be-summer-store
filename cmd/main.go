package main


func main() {
	cfg := config.LoadConfig()
	database.InitDB(cfg.DatabaseDSN)

	//database.DB.AutoMigrate(&model.Product{})

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/products", controller.GetProducts)
	}

	r.Run(":" + cfg.Port)
}
