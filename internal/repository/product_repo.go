package repository

import (
	"stbe/internal/database"
	"stbe/internal/model"
)

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := database.DB.Where("status = ?", "active").Order("created_at desc").Find(&products).Error
	return products, err
}
