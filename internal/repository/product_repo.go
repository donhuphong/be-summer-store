package repository

import (
	"be-summer-store/internal/database"
	"be-summer-store/internal/model"
)

func GetAllProducts() ([]model.Product, error) {
	var products []model.Product

	err := database.DB.
		Where("status = ?", "active").
		Order("created_at desc").
		Find(&products).Error

	return products, err
}

func GetProductByID(id uint) (*model.Product, error) {
	var product model.Product

	err := database.DB.
		First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func GetProducts(page int, limit int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	offset := (page - 1) * limit

	// count total
	if err := database.DB.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// get data
	err := database.DB.
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func InsertProduct(product *model.Product) error {
	return database.DB.Create(product).Error
}

func UpdateProduct(product *model.Product) error {
	return database.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&model.Product{}, id).Error
}
