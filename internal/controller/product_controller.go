package controller

import (
	"be-summer-store/internal/model"
	"be-summer-store/internal/service"
	"net/http"
	"time"

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

func CreateProduct(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	err := service.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "product created",
		"data":    product,
	})
}
