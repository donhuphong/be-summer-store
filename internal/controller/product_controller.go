package controller

import (
	"net/http"
	"stbe/internal/service"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := service.GetActiveProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi xử lý dữ liệu sản phẩm"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}
