package utilities

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString() string {
	var length int = 6
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
