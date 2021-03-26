package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisBackend struct {
	ctx       context.Context
	client    *redis.Client
	prefixVal string
	prefixHsh string
}

func (b *RedisBackend) Get(key string) (res string, err error) {
	return b.get(b.prefixVal, key)
}

func (b *RedisBackend) GetHash(key string) (res string, err error) {
	return b.get(b.prefixHsh, key)
}

func (b *RedisBackend) Set(key, value string, ttl time.Duration) (err error) {
	return b.set(b.prefixVal, key, value, ttl)
}

func (b *RedisBackend) SetHash(key, value string, ttl time.Duration) (err error) {
	return b.set(b.prefixHsh, key, value, ttl)
}

///

func (b *RedisBackend) get(prefix, key string) (res string, err error) {
	cmd := b.client.Get(b.ctx, prefix+key)
	if err := cmd.Err(); err != nil && err != redis.Nil {
		return "", err
	}
	res, _ = cmd.Result()
	return
}

func (b *RedisBackend) set(prefix, key, value string, ttl time.Duration) (err error) {
	cmd := b.client.Set(b.ctx, prefix+key, value, ttl)
	err = cmd.Err()
	return
}
