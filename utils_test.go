package utilities

import (
	"errors"
	"testing"
)

func TestFunc(t *testing.T) {
	FormatError("Asia/Shanghai", "2006-01-02 15:04:05", errors.New("damn error"))
}

func TestEnDe(t *testing.T) {
	var err error
	var username string = "adamliuio"
	var key string = "adamadamadamadam"
	var hash string = "jn+UBtmhxOOUFwHBirAUiTpvIuRskHp6tw"
	for i := 0; i < 10; i++ {
		if hash, err = Encrypt(username, key); err != nil { // turn encrypted cookie into username
			t.Fatal(err)
		}
		t.Log("hash:", hash)
		if username, err = Decrypt(hash, key); err != nil { // turn encrypted cookie into username
			t.Fatal(err)
		}
		t.Log("username:", username)
	}
}
