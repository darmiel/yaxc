package server

import (
	"time"
	"zgo.at/zcache"
)

type CacheBackend struct {
	cache   *zcache.Cache
	errCast error
}

func (b *CacheBackend) Get(key string) (res string, err error) {
	return b.get("val::" + key)
}

func (b *CacheBackend) GetHash(key string) (res string, err error) {
	return b.get("hash::" + key)
}

func (b *CacheBackend) Set(key, value string, ttl time.Duration) error {
	b.cache.Set("val::"+key, value, ttl)
	return nil
}

func (b *CacheBackend) SetHash(key, value string, ttl time.Duration) error {
	b.cache.Set("hash::"+key, value, ttl)
	return nil
}

func (b *CacheBackend) get(key string) (res string, err error) {
	if v, ok := b.cache.Get(key); ok {
		if s, ok := v.(string); ok {
			return s, nil
		}
		return "", b.errCast
	}
	return "", nil
}
