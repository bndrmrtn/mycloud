package permissions

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()

func redisBoolReturn(rdb *redis.Client, key string) (bool, error) {
	res, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return res == "true", nil
}

func redisBoolSave(rdb *redis.Client, val bool, key string) bool {
	err := rdb.Set(ctx, key, fmt.Sprintf("%v", val), time.Hour*1).Err()
	if err != nil {
		logrus.Warn("Failed to save to redis: ", err)
	}
	return val
}
