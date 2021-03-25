package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"math"
)

var (
	errKeyEmpty    = errors.New("key empty")
	errKeyLength   = errors.New("invalid key length")
	errNonceLength = errors.New("invalid nonce length")
	bytesEmpty     []byte
)

func Encrypt(textStr, keyStr string) (res []byte, err error) {
	gcm, err := getGCM(keyStr)
	if err != nil {
		return bytesEmpty, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return bytesEmpty, err
	}
	seal := gcm.Seal(nonce, nonce, []byte(textStr), nil)
	res = []byte(base64.StdEncoding.EncodeToString(seal))
	return
}

func Decrypt(textStr, keyStr string) (res []byte, err error) {
	text, err := base64.StdEncoding.DecodeString(textStr)
	if err != nil {
		return bytesEmpty, err
	}
	gcm, err := getGCM(keyStr)
	if err != nil {
		return bytesEmpty, err
	}
	nonceSize := gcm.NonceSize()
	if len(text) < nonceSize {
		return bytesEmpty, errNonceLength
	}
	nonce, encrypted := text[:nonceSize], text[nonceSize:]
	plain, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return bytesEmpty, err
	}
	return plain, nil
}

////

func getGCM(keyStr string) (cipher.AEAD, error) {
	key, err := getKey(keyStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

func getKey(keyStr string) (key []byte, err error) {
	key = []byte(keyStr)
	if len(key) <= 0 {
		return key, errKeyEmpty
	}
	if len(key) < 32 {
		rep := math.Ceil(32.0 / float64(len(keyStr)))
		key = bytes.Repeat([]byte(keyStr), int(rep))
	}
	if len(key) > 32 {
		key = key[:32]
	}
	if len(key) != 32 {
		return key, errKeyLength
	}
	return
}
