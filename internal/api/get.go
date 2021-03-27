package api

import (
	"github.com/darmiel/yaxc/internal/common"
	"github.com/imroc/req"
)

func (a *api) GetContent(path, passphrase string) (res string, err error) {
	var resp *req.Resp
	if resp, err = req.Get(a.ServerURL + "/" + path); err != nil {
		return
	}
	res = resp.String()
	// encryption
	if passphrase != "" {
		var b []byte
		if b, err = common.Decrypt(res, passphrase); err != nil {
			return
		}
		res = string(b)
	}
	return
}

func (a *api) SetContent(path, passphrase, content string) (err error) {
	if passphrase != "" {
		var b []byte
		if b, err = common.Encrypt(content, passphrase); err != nil {
			return
		}
		content = string(b)
	}
	_, err = req.Post(a.ServerURL+"/"+path, content)
	return
}
