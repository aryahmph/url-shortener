package cache

import (
	"context"
	"time"
)

//go:generate mockery --name=Repository
type Repository interface {
	Set(ctx context.Context, key, value string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
