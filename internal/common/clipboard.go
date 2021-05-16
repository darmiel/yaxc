// +build client

package common

import (
	"encoding/base64"
	"github.com/atotto/clipboard"
)

func GetClipboard(b64 bool) (res string, err error) {
	if res, err = clipboard.ReadAll(); err != nil {
		return
	}
	// encode base64
	if b64 && res != "" {
		res = base64.StdEncoding.EncodeToString([]byte(res))
	}
	return
}

func WriteClipboard(txt string, b64 bool) (err error) {
	// decode base64
	if b64 {
		var b []byte
		if b, err = base64.StdEncoding.DecodeString(txt); err != nil {
			return
		}
		txt = string(b)
	}
	err = clipboard.WriteAll(txt)
	return
}
