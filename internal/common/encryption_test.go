package common

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	enc, err := Encrypt("hello", "world")
	if err != nil {
		t.Error(err)
		return
	}
	str := string(enc)
	log.Println("Encrypted:", str)
	if str == "hello" {
		t.Errorf("encrypted string equal to input")
		return
	}
	dec, err := Decrypt(str, "world")
	if err != nil {
		t.Error(err)
		return
	}
	str = string(dec)
	log.Println("Decrypted:", str)
	AssertEqual(t, str, "hello")
}

func TestDecrypt(t *testing.T) {
	strin := "Z6wyotQ5w/9dMvZnjMzGciP6p8+zvcOVr6tuKhGYRKbx"
	decrypt, err := Decrypt(strin, "world")
	if err != nil {
		t.Error(err)
		return
	}
	AssertEqual(t, string(decrypt), "hello")
}
