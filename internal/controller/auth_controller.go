package controller

import (
	"be-summer-store/internal/config"
	"be-summer-store/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login
func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Fake user validation
	if req.Username != "admin" || req.Password != config.AppConfig.JWT.PASS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	userID := uint(1)

	accessToken, err := utils.GenerateAccessToken(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid GenerateAccessToken",
		})
		return
	}
	refreshToken, err := utils.GenerateRefreshToken(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid GenerateRefreshToken",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Refresh Access Token
func RefreshToken(c *gin.Context) {

	type Request struct {
		RefreshToken string `json:"refresh_token"`
	}

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims, err := utils.ParseRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid refresh token",
		})
		return
	}

	newAccessToken, _ := utils.GenerateAccessToken(claims.UserID)

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
