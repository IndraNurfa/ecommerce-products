package repository

import (
	"context"
	"ecommerce-products/internal/models"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func (r *CategoryRepo) InsertNewCategory(ctx context.Context, category *models.ProductCategory) error {
	return r.DB.Create(category).Error
}
