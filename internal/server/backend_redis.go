package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisBackend struct {
	ctx    context.Context
	client *redis.Client
	prefix string
}

func (b *RedisBackend) Get(key string) (res string, err error) {
	cmd := b.client.Get(b.ctx, b.prefix+key)
	if err := cmd.Err(); err != nil && err != redis.Nil {
		return "", err
	}
	res, _ = cmd.Result()
	return
}

func (b *RedisBackend) Set(key, value string, ttl time.Duration) (err error) {
	cmd := b.client.Set(b.ctx, b.prefix+key, value, ttl)
	err = cmd.Err()
	return
}
