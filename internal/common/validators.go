package common

import (
	"log"
	"regexp"
)

const (
	APMinLength = 3
	APMaxLength = 128
)

var (
	APAllowedChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.-_+*!$%~@")
)

func ValidateAnywherePath(anywhere string) bool {
	l := len(anywhere)
	if l < APMinLength || l > APMaxLength {
		return false
	}
	for _, c := range anywhere {
		b := byte(c)
		f := false
		for _, a := range APAllowedChars {
			if a == b {
				f = true
				break
			}
		}
		if !f {
			return false
		}
	}
	return true
}

var p *regexp.Regexp

func init() {
	var err error
	if p, err = regexp.Compile(`^[A-Za-z0-9.\-_+*!$%~@]{3,128}$`); err != nil {
		log.Fatalln(err)
	}
}

func ValidateAnywherePathRegex(anywhere string) bool {
	return p.MatchString(anywhere)
}
