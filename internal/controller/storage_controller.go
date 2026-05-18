package controller

import (
	"be-summer-store/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/presign?key=products/anh.jpg
func GetPresignURL(c *gin.Context) {
	key := c.Query("key")

	result, err := service.GetPresignURL(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": result})
}
