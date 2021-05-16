// +build client

package api

import (
	"errors"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/imroc/req"
)

var (
	ErrErrResponse = errors.New("invalid response")
)

func (a *Api) GetContent(path, passphrase string) (res string, err error) {
	var resp *req.Resp
	if resp, err = req.Get(a.UrlGetContents(path, "")); err != nil {
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

func (a *Api) SetContent(path, passphrase, content string) (err error) {
	if passphrase != "" {
		var b []byte
		if b, err = common.Encrypt(content, passphrase); err != nil {
			return
		}
		content = string(b)
	}
	hash := common.MD5Hash(content)
	// with custom hash
	url := a.UrlSetContents(path, hash, "", 0)
	_, err = req.Post(url, content)
	return
}

func (a *Api) GetHash(path string) (res string, err error) {
	var resp *req.Resp
	if resp, err = req.Get(a.UrlGetHash(path)); err != nil {
		return
	}

	if resp.Response().StatusCode != 200 {
		err = ErrErrResponse
		return
	}

	res = resp.String()
	return
}
