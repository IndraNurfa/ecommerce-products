package services

import (
	"context"
	"ecommerce-products/internal/interfaces"
	"ecommerce-products/internal/models"

	"github.com/pkg/errors"
)

type ProductService struct {
	ProductRepo interfaces.IProductRepo
}

// func (s *ProductService) (ctx context.Context)
func (s *ProductService) CreateProduct(ctx context.Context, req *models.Product) (*models.Product, error) {
	err := s.ProductRepo.InsertNewProduct(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert new product")
	}

	resp := req

	return resp, nil
}
