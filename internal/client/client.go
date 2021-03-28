package client

import (
	"errors"
	"github.com/darmiel/yaxc/internal/api"
	"sync"
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
}

func NewCheck(path, pass string) *Check {
	return &Check{
		a:    api.API(),
		mu:   sync.Mutex{},
		path: path,
		pass: pass,
	}
}
