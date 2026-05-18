package router

import (
	middleware "be-summer-store/internal/auth"
	"be-summer-store/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	// Public routes
	{
		api.POST("/login", controller.Login)
		api.POST("/refresh", controller.RefreshToken)

		api.GET("/products", controller.GetProducts)
	}

	// Private
	protected := api.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.POST("/products", controller.CreateProduct)
		protected.GET("/presign", controller.GetPresignURL)
	}

	return r
}
