package client

import (
	"github.com/atotto/clipboard"
	"github.com/darmiel/yaxc/internal/common"
	"log"
	"strings"
	"time"
)

func WatchClient(c *Check, d time.Duration, done chan bool) {
	t := time.NewTicker(d)
	for {
		select {
		case <-done:
			log.Println("[Watch Client] Stopping ...")
			return
		case <-t.C:
			if err := c.CheckClient(); err != nil {
				log.Println("[Watch Client] WARN:", err)
			}
			break
		}
	}
}

func (c *Check) CheckClient() (err error) {
	// check if clipboard actions are available
	if clipboard.Unsupported {
		return ErrUnsupported
	}

	// lock
	c.mu.Lock()
	defer c.mu.Unlock()

	// get clipboard
	var cb string
	cb, _ = clipboard.ReadAll()

	// ignore empty clipboard
	if strings.TrimSpace(cb) == "" {
		return
	}

	// calculate hash
	var ch string
	if ch = common.MD5Hash(cb); ch == "" {
		err = ErrEmptyHash
		return
	}

	// get hash from server
	var sh string
	sh, _ = c.a.GetHash(c.path)
	if strings.TrimSpace(sh) != "" {
		// compare hashes
		if ch == sh {
			return
		}
	}

	// upload to server
	err = c.a.SetContent(c.path, c.pass, cb)
	log.Println("Wrote: '" + cb + "' to server")
	return
}
