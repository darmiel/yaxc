package client

import (
	"errors"
	"github.com/darmiel/yaxc/internal/api"
	"sync"
	"time"
)

var (
	ErrEmptyHash   = errors.New("empty hash")
	ErrUnsupported = errors.New("clipboard unsupported")
)

type Check struct {
	a    *api.Api
	mu   sync.Mutex
	path string
	pass string
	//
	previousClipboard string
	lastUpdate        time.Time
}

func NewCheck(path, pass string) *Check {
	return &Check{
		a:    api.API(),
		mu:   sync.Mutex{},
		path: path,
		pass: pass,
	}
}
