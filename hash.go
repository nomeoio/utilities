package utilities

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GenerateHash() string {
	var s string = fmt.Sprint(time.Now().UnixNano())
	var h [32]byte = sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", h)
}
