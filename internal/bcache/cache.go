package bcache

import (
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/muesli/termenv"
	"sync"
	"time"
)

var prefix termenv.Style

func init() {
	p := common.Profile()
	prefix = termenv.String("CCHE").Foreground(p.Color("0")).Background(p.Color("#D290E4"))
}

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

	// TODO: remove debug
	fmt.Println(prefix,
		termenv.String("<-").Foreground(common.Profile().Color("#DBAB79")),
		"Set",
		termenv.String(key).Foreground(common.Profile().Color("#A8CC8C")),
		termenv.String("=").Foreground(common.Profile().Color("#DBAB79")),
		value)

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

			// TODO: remove debug
			fmt.Println(prefix,
				termenv.String("->").Foreground(common.Profile().Color("#66C2CD")),
				"Get",
				termenv.String(key).Foreground(common.Profile().Color("#A8CC8C")),
				termenv.String("=").Foreground(common.Profile().Color("#DBAB79")),
				v.value)

			c.mu.Unlock()
			return v.value, true
		}
	}
	c.mu.Unlock()
	return nil, false
}
