// +build client

package client

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/darmiel/yaxc/internal/common"
	"time"
)

func WatchClient(c *Check, d time.Duration, done chan bool) {
	t := time.NewTicker(d)
	for {
		select {
		case <-done:
			fmt.Println(common.StyleInfo(), "Stopping", common.WordClient(), "Watcher")
			return
		case <-t.C:
			if err := c.CheckClient(); err != nil {
				fmt.Println(common.StyleWarn(), err)
			}
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
	cb, _ = common.GetClipboard(c.useBase64)

	// ignore empty clipboard
	if cb == "" {
		return
	}

	if c.previousClipboard == cb {
		return
	}
	c.previousClipboard = cb

	// upload to server
	err = c.a.SetContent(c.path, c.pass, cb)
	fmt.Println(common.StyleUpdate(), common.WordServer(), "<-", common.PrettyLimit(cb, 48))
	return
}
