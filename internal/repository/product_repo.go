package repository

import (
	"be-summer-store/internal/database"
	"be-summer-store/internal/model"
)

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := database.DB.Where("status = ?", "active").Order("created_at desc").Find(&products).Error
	return products, err
}
