package repository

import (
	"context"
	"ecommerce-products/constants"
	"ecommerce-products/helpers"
	"ecommerce-products/internal/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB    *gorm.DB
	Redis *redis.ClusterClient
}

// func(r *ProductRepo) (ctx context.Context, )

func (r *ProductRepo) InsertNewProduct(ctx context.Context, product *models.Product) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {

		err := tx.Create(product).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err == nil {
		go func() {
			jsonData, err := json.Marshal(product)
			if err != nil {
				helpers.Logger.Warn("failed to marshal data product")
				return
			}

			if err := r.Redis.Del(ctx, constants.RedisKeyProducts).Err(); err != nil {
				helpers.Logger.Warn("failed to delete redis key: ", constants.RedisKeyProducts)
			}
			if err := r.Redis.Del(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, product.ID)).Err(); err != nil {
				helpers.Logger.Warn("failed to delete redis key: ", fmt.Sprintf(constants.RedisKeyProductDetail, product.ID))
			}
			if err := r.Redis.Set(ctx, fmt.Sprintf(constants.RedisKeyProductDetail, product.ID), string(jsonData), time.Hour*24).Err(); err != nil {
				helpers.Logger.Warn("failed to set redis key: ", fmt.Sprintf(constants.RedisKeyProductDetail, product.ID))
			}
		}()
	}

	return err
}
