package service

import (
	"be-summer-store/internal/model"
	"be-summer-store/internal/repository"
)

func GetActiveProducts() ([]model.Product, error) {
	products, err := repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var filteredProducts []model.Product
	for _, p := range products {
		if p.Stock > 0 {
			filteredProducts = append(filteredProducts, p)
		}
	}
	return filteredProducts, nil
}
