package interfaces

import (
	"context"
	"ecommerce-products/external"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (external.Profile, error)
}
