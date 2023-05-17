package redis

import (
	"context"
	"errors"
	cacheRepo "github.com/aryahmph/url-shortener/internal/repository/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisCacheRepository struct {
	client *redis.Client
}

func NewRedisCacheRepository(client *redis.Client) cacheRepo.Repository {
	return &redisCacheRepository{client: client}
}

func (r redisCacheRepository) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return r.client.SetEx(ctx, key, value, expiration).Err()
}

func (r redisCacheRepository) Get(ctx context.Context, key string) (string, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("cache key not found")
	} else if err != nil {
		return "", err
	}
	return res, nil
}
