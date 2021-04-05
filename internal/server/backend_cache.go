package server

import (
	"github.com/darmiel/yaxc/internal/bcache"
	"time"
)

type CacheBackend struct {
	cache   *bcache.Cache
	errCast error
}

func (b *CacheBackend) Get(key string) (res string, err error) {
	return b.getString("val::" + key)
}

func (b *CacheBackend) GetHash(key string) (res string, err error) {
	return b.getString("hash::" + key)
}

func (b *CacheBackend) Set(key, value string, ttl time.Duration) error {
	b.cache.Set("val::"+key, []byte(value), ttl)
	return nil
}

func (b *CacheBackend) SetHash(key, value string, ttl time.Duration) error {
	b.cache.Set("hash::"+key, []byte(value), ttl)
	return nil
}

func (b *CacheBackend) getString(key string) (res string, err error) {
	if v, ok := b.cache.Get(key); ok {
		if r, o := v.([]byte); o {
			res = string(r)
		} else {
			err = b.errCast
		}
	}
	return
}
