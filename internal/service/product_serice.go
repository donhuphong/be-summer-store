package service

import (
	"stbe/internal/model"
	"stbe/internal/repository"
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
