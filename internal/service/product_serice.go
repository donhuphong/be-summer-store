package service

import (
	"be-summer-store/internal/model"
	"be-summer-store/internal/repository"
	"errors"
	"math"
)

type ProductPagingResponse struct {
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
	Total      int64           `json:"total"`
	TotalPages int             `json:"total_pages"`
	Data       []model.Product `json:"data"`
}

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
func GetProductPaging(page int, limit int) (*ProductPagingResponse, error) {
	products, total, err := repository.GetProducts(page, limit)
	if err != nil {
		return nil, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	response := &ProductPagingResponse{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Data:       products,
	}
	return response, nil
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

func UpdateProduct(product *model.Product) error {
	if product.Name == "" {
		return errors.New("name is required")
	}
	if product.Price <= 0 {
		return errors.New("invalid price")
	}
	return repository.UpdateProduct(product)
}
