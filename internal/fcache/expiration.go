package fcache

import "time"

// EXPIRATION

const ExpirationDefault time.Duration = -1
const ExpirationKeep time.Duration = 0

type nodeExpiration int64

func (e nodeExpiration) IsExpired() bool {
	if e == 0 {
		return false
	}
	return time.Now().Unix() > int64(e)
}

func (c *Cache) expiration(d time.Duration) nodeExpiration {
	if d == ExpirationKeep {
		return 0
	}
	if d == ExpirationDefault {
		d = c.de
	}
	return nodeExpiration(time.Now().Add(d).Unix())
}
