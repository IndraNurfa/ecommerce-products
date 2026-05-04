package helpers

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.ClusterClient

func SetupRedis() {

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(GetEnv("REDIS_HOST", "localhost:6379"), ","),
	})

	ctx := context.Background()

	// Better than Ping for cluster validation
	err := client.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})

	if err != nil {
		Logger.Error("Failed to connect redis: ", err)
		return
	}
	Logger.Info("Redis cluster connected successfully")

	RedisClient = client
}
