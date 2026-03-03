package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
)

// RedisConfig represents the configuration for a Redis database connection
type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// RedisClient represents a Redis client connection
type RedisClient struct {
	*redis.Client
}

// NewRedisClient returns a new Redis client instance
func NewRedisClient(cfg RedisConfig) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,
		DialTimeout: 10 * time.Second,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{client}, nil
}

// GetRedisClient returns a Redis client instance from the configuration
func GetRedisClient() (*RedisClient, error) {
	cfg := RedisConfig{
		Address: "localhost:6379",
		Password: "",
		DB:       0,
	}

	return NewRedisClient(cfg)
}

// SetKey sets a key-value pair in Redis
func (r *RedisClient) SetKey(ctx context.Context, key string, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

// GetKey retrieves a value from Redis by key
func (r *RedisClient) GetKey(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil && errors.Is(err, redis.Nil) {
		return "", nil
	}
	return val, err
}

// DeleteKey deletes a key from Redis
func (r *RedisClient) DeleteKey(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
