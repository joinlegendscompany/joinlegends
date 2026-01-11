package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func New() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func Ctx() context.Context {
	return context.Background()
}
