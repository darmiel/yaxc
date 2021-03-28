package client

import (
	"github.com/atotto/clipboard"
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

	if c.previousClipboard == cb {
		return
	}
	c.previousClipboard = cb

	// upload to server
	err = c.a.SetContent(c.path, c.pass, cb)
	log.Println("Wrote: '" + cb + "' to server")
	return
}
