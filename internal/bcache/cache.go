package bcache

import (
	"sync"
	"time"
)

type node struct {
	expires nodeExpiration
	value   interface{}
}

type Cache struct {
	mu                sync.Mutex
	values            map[string]*node
	defaultExpiration time.Duration
	cleanerInterval   time.Duration
}

func NewCache(defaultExpiration, cleanerInterval time.Duration) *Cache {
	c := &Cache{
		values:            make(map[string]*node),
		defaultExpiration: defaultExpiration,
		cleanerInterval:   cleanerInterval,
	}
	if cleanerInterval != 0 {
		// go c.janitorService()
	}
	return c
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()
	printDebugSet(key, value)
	c.values[key] = &node{
		expires: c.expiration(expiration),
		value:   value,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	if v, o := c.values[key]; o && v != nil {
		if !v.expires.IsExpired() {
			printDebugGet(key, v.value)
			c.mu.Unlock()
			return v.value, true
		}
	}
	c.mu.Unlock()
	return nil, false
}
