package fcache

import (
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/muesli/termenv"
	"time"
)

func (c *Cache) janitorService() {
	if c.cleanerInterval == 0 {
		return
	}
	for {
		time.Sleep(c.cleanerInterval)
		fmt.Println(prefix,
			termenv.String("JANITOR").Foreground(common.Profile().Color("#A8CC8C")),
			"Starting ...")
		c.janitor()
	}
}

func (c *Cache) janitor() {
	c.mu.Lock()
	for k, v := range c.val {
		// nil node
		if v == nil || v.expires.IsExpired() {
			fmt.Println(prefix,
				termenv.String("JANITOR").Foreground(common.Profile().Color("#A8CC8C")),
				"Deleting", termenv.String(k).Foreground(common.Profile().Color("#A8CC8C")))
			delete(c.val, k)
		}
	}
	c.mu.Unlock()
}
