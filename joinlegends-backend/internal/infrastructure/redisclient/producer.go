package redisclient

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func Send(ctx context.Context, rdb *redis.Client, msg QueueMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return rdb.LPush(
		ctx,
		ViewQueue,
		data,
	).Err()
}
