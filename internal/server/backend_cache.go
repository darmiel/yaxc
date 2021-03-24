package server

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type CacheBackend struct {
	c       *cache.Cache
	errCast error
}

func (b *CacheBackend) Get(key string) (res string, err error) {
	if v, ok := b.c.Get(key); ok {
		if s, ok := v.(string); ok {
			return s, nil
		}
		return "", b.errCast
	}
	return "", nil
}

func (b *CacheBackend) Set(key, value string, ttl time.Duration) error {
	b.c.Set(key, value, ttl)
	return nil
}
