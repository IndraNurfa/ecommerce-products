package interfaces

import (
	"context"
	"ecommerce-products/internal/models"
)

type IHealthService interface {
	CheckHealthConnection(ctx context.Context) (*models.Health, error)
}

type IHealthRepository interface {
	CheckDatabaseConnection(ctx context.Context) error
}
