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
	//
	useBase64 bool
	//
	previousClipboard string
}

func NewCheck(path, pass string, b64 bool) *Check {
	return &Check{
		a:         api.API(),
		mu:        sync.Mutex{},
		path:      path,
		pass:      pass,
		useBase64: b64,
	}
}
