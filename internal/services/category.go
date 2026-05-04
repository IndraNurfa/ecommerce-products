package services

import (
	"context"
	"ecommerce-products/internal/interfaces"
	"ecommerce-products/internal/models"

	"github.com/pkg/errors"
)

type CategoryService struct {
	CategoryRepo interfaces.ICategoryRepo
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *models.ProductCategory) (*models.ProductCategory, error) {
	err := s.CategoryRepo.InsertNewCategory(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert new product")
	}
	resp := req
	return resp, nil
}
