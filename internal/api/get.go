package api

import (
	"errors"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/imroc/req"
)

var (
	ErrErrResponse = errors.New("invalid response")
)

func (a *api) GetContent(path, passphrase string) (res string, err error) {
	var resp *req.Resp
	if resp, err = req.Get(a.ServerURL + "/" + path); err != nil {
		return
	}

	if resp.Response().StatusCode != 200 {
		err = ErrErrResponse
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

func (a *api) GetHash(path string) (res string, err error) {
	var resp *req.Resp
	if resp, err = req.Get(a.ServerURL + "/hash/" + path); err != nil {
		return
	}

	if resp.Response().StatusCode != 200 {
		err = ErrErrResponse
		return
	}

	res = resp.String()
	return
}
