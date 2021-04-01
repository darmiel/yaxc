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
	APAllowedChars  = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789:.-_+*!$%~@")
	HexAllowedChars = []byte("0123456789abcdef")
)

func ContainsOtherThan(str string, allowed []byte) bool {
	for _, c := range str {
		b := byte(c)
		v := false
		// check if allowed array contains
		for _, a := range allowed {
			if b == a {
				v = true
				break
			}
		}
		if !v {
			return true
		}
	}
	return false
}

func ValidateAnywherePath(anywhere string) bool {
	l := len(anywhere)
	if l < APMinLength || l > APMaxLength {
		return false
	}
	return !ContainsOtherThan(anywhere, APAllowedChars)
}

func ValidateHex(anywhere string) bool {
	l := len(anywhere)
	if l < 1 || l > 256 {
		return false
	}
	return !ContainsOtherThan(anywhere, HexAllowedChars)
}

var p *regexp.Regexp

func init() {
	var err error
	if p, err = regexp.Compile(`^[A-Za-z0-9.\-:_+*!$%~@]{3,128}$`); err != nil {
		log.Fatalln(err)
	}
}

func ValidateAnywherePathRegex(anywhere string) bool {
	return p.MatchString(anywhere)
}
