package server

import (
	"time"
)

type Backend interface {
	Get(key string) (string, error)
	Set(key, value string, ttl time.Duration) error

	GetHash(key string) (string, error)
	SetHash(key, value string, ttl time.Duration) error
}
