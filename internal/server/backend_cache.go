package server

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type CacheBackend struct {
	cacheVal *cache.Cache
	cacheHsh *cache.Cache
	errCast  error
}

func (b *CacheBackend) Get(key string) (res string, err error) {
	return b.get(b.cacheVal, key)
}

func (b *CacheBackend) GetHash(key string) (res string, err error) {
	return b.get(b.cacheHsh, key)
}

func (b *CacheBackend) Set(key, value string, ttl time.Duration) error {
	log.Info("Updating cache cacheVal with key", key, "and ttl", ttl)
	b.cacheVal.Set(key, value, ttl)
	return nil
}

func (b *CacheBackend) SetHash(key, value string, ttl time.Duration) error {
	log.Info("Updating cache cacheHsh with key", key, "and ttl", ttl)
	b.cacheHsh.Set(key, value, ttl)
	return nil
}

func (b *CacheBackend) get(c *cache.Cache, key string) (res string, err error) {
	var cName string
	if c == b.cacheVal {
		cName = "val"
	} else if c == b.cacheHsh {
		cName = "hsh"
	} else {
		cName = "unknown"
	}

	r1, r2, r3 := c.GetWithExpiration(key)
	log.Info("Requesting cache >", cName, "(", c, ")", "< with key", key, "with result:", r1, "=>", r2, "=>", r3)

	if v, ok := c.Get(key); ok {
		if s, ok := v.(string); ok {
			return s, nil
		}
		return "", b.errCast
	}
	return "", nil
}
