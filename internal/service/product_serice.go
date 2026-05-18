package service

import (
	"be-summer-store/internal/model"
	"be-summer-store/internal/repository"
	"errors"
)

func GetActiveProducts() ([]model.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var filteredProducts []model.Product
	for _, p := range products {
		if p.Stock > -1 {
			filteredProducts = append(filteredProducts, p)
		}
	}
	return filteredProducts, nil
}

func CreateProduct(product *model.Product) error {
	if product.Name == "" {
		return errors.New("name is required")
	}

	if product.Price <= 0 {
		return errors.New("invalid price")
	}
	return repository.InsertProduct(product)
}
