package utilities

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
)

func RandomString() string {
	var length int = 8
	b := make([]byte, length)
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	math_rand.Read(b)
	return fmt.Sprintf("%x", b)
}
