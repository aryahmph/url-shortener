package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func New(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return client
}
