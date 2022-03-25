package utilities

import "testing"

func TestRandomString(t *testing.T) {
	for i := 1; i < 25; i++ {
		t.Log(RandomString())
	}
}
