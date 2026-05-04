package interfaces

import (
	"context"
	"ecommerce-products/internal/models"

	"github.com/labstack/echo/v5"
)

type ICategoryAPI interface {
	CreateCategory(e *echo.Context) error
}

type ICategoryService interface {
	CreateCategory(ctx context.Context, req *models.ProductCategory) (*models.ProductCategory, error)
}

type ICategoryRepo interface {
	InsertNewCategory(ctx context.Context, category *models.ProductCategory) error
}
