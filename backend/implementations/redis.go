package implementations

import (
	"context"
	"time"

	"github.com/bndrmrtn/go-gale"
	"github.com/redis/go-redis/v9"
)

type RedisSessionStore struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedisSessionStore(ctx context.Context, client *redis.Client) gale.SessionStore {
	return &RedisSessionStore{
		ctx:    ctx,
		client: client,
	}
}

func (r *RedisSessionStore) Get(key string) ([]byte, error) {
	s, err := r.client.Get(r.ctx, key).Result()
	return []byte(s), err
}

func (r *RedisSessionStore) Exists(key string) bool {
	ok, err := r.client.Exists(r.ctx, key).Result()
	return ok > 0 && err == nil
}

func (r *RedisSessionStore) Set(key string, val []byte) error {
	return r.SetEx(key, val, 0)
}

func (r *RedisSessionStore) SetEx(key string, val []byte, expiry time.Duration) error {
	return r.client.Set(r.ctx, key, val, expiry).Err()
}

func (r *RedisSessionStore) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
