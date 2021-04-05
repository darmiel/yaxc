package bcache

import (
	"time"
)

func (c *Cache) janitorService() {
	if c.cleanerInterval == 0 {
		return
	}
	for {
		time.Sleep(c.cleanerInterval)
		printDebugJanitorStart()
		c.janitor()
	}
}

func (c *Cache) janitor() {
	c.mu.Lock()
	for k, v := range c.values {
		// nil node
		if v == nil || v.expires.IsExpired() {
			printDebugJanitorDelete(k)
			delete(c.values, k)
		}
	}
	c.mu.Unlock()
}
