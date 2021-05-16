// +build client

package api

import (
	"strings"
	"time"
)

func (a *Api) urlBuilder(anywhere string) (b *strings.Builder) {
	b = new(strings.Builder)
	b.WriteString(a.ServerURL)
	b.WriteRune('/')
	b.WriteString(anywhere)
	return
}

func (a *Api) UrlSetContents(anywhere, hash, secret string, ttl time.Duration) string {
	b := a.urlBuilder(anywhere)
	// hash
	if hash != "" {
		b.WriteByte('/')
		b.WriteString(hash)
	}
	// secret
	if secret != "" {
		b.WriteString("?secret=")
		b.WriteString(secret)
	}
	// ttl
	if ttl != 0 {
		if secret == "" {
			b.WriteByte('?')
		} else {
			b.WriteByte('&')
		}
		b.WriteString("ttl=")
		b.WriteString(ttl.String())
	}
	return b.String()
}

func (a *Api) UrlGetContents(anywhere, secret string) string {
	b := a.urlBuilder(anywhere)
	if secret != "" {
		b.WriteString("?secret=")
		b.WriteString(secret)
	}
	return b.String()
}

func (a *Api) UrlGetHash(anywhere string) string {
	var b strings.Builder
	b.WriteString(a.ServerURL)
	b.WriteString("/hash/")
	b.WriteString(anywhere)
	return b.String()
}
