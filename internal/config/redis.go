package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect redis: %w", err)
	}

	Redis = client

	return nil
}

func GetRedis() *redis.Client {
	return Redis
}
