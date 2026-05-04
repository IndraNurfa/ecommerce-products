package interfaces

import (
	"context"
	"ecommerce-products/internal/models"

	"github.com/labstack/echo/v5"
)

type IProductAPI interface {
	CreateProduct(e *echo.Context) error
}

type IProductService interface {
	CreateProduct(ctx context.Context, req *models.Product) (*models.Product, error)
}

type IProductRepo interface {
	InsertNewProduct(ctx context.Context, product *models.Product) error
}
