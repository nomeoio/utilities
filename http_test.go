package utilities

import (
	"testing"
)

func TestRedirect(t *testing.T) {
	var err error
	var url string
	if url, err = GetRedirectedLink("https://bit.ly/AdamLiuYoutube"); err != nil {
		t.Log(err)
	}
	t.Log(url)
}
