package client

import (
	"github.com/atotto/clipboard"
	"github.com/darmiel/yaxc/internal/api"
	"github.com/darmiel/yaxc/internal/common"
	"log"
	"time"
)

func WatchServer(c *Check, d time.Duration, done chan bool) {
	t := time.NewTicker(d)
	for {
		select {
		case <-done:
			log.Println("[Watch Server] Stopping ...")
			return
		case <-t.C:
			if err := c.CheckServer(); err != nil {
				log.Println("[Watch Server] WARN:", err)
			}
			break
		}
	}
}

func (c *Check) CheckServer() (err error) {
	// check if clipboard actions are available
	if clipboard.Unsupported {
		return ErrUnsupported
	}

	// lock
	c.mu.Lock()
	defer c.mu.Unlock()

	// get hash from server
	var sh string
	if sh, err = c.a.GetHash(c.path); err != nil {
		if err == api.ErrErrResponse {
			err = nil
		}
		return
	}
	if sh == "" {
		return ErrEmptyHash
	}

	// get clipboard
	var cb string
	cb, _ = clipboard.ReadAll()
	if cb != "" {
		// calculate hash
		var ch string
		if ch = common.MD5Hash(cb); ch == "" {
			err = ErrEmptyHash
			return
		}

		// compare hashes
		if ch == sh {
			return
		}
	}

	// get data from server
	var sd string
	if sd, err = c.a.GetContent(c.path, c.pass); err != nil {
		return
	}

	// update contents
	err = clipboard.WriteAll(sd)
	log.Println("Wrote: '" + sd + "' to client")
	return
}
