package common

import (
	"crypto/md5"
	"fmt"
)

func Hash(text string) string {
	data := []byte(text)
	sum := md5.Sum(data)
	return fmt.Sprintf("%x", sum)
}
