package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)
// GET /api/presign?key=products/anh.jpg
func GetPresignURL(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thiếu tham số key"})
		return
	}

	presignClient := s3.NewPresignClient(r2Client)
	req, err := presignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(10*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo presigned URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"url":        req.URL,
			"public_url": publicBaseURL + "/" + key,
		},
	})
}

// POST /api/products
type CreateProductRequest struct {
	Name        string  `json:"name"        binding:"required"`
	Price       float64 `json:"price"       binding:"required,gt=0"`
	Description string  `json:"description"`
	CategoryID  int64   `json:"category_id" binding:"required"`
	Stock       int     `json:"stock"       binding:"min=0"`
	ImageURL    string  `json:"image_url"   binding:"required,url"`
}

func CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := service.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi tạo sản phẩm"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   product,
	})
}

