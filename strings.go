package utilities

import (
	"net/http"
	"net/mail"
)

func CheckUrl(url string) (finalUrl string, contentLength int64, err error) {
	// check redirected final url & remove file size
	var resp *http.Response
	if resp, err = http.Head(url); err != nil {
		return
	}

	finalUrl = resp.Request.URL.String()
	contentLength = resp.ContentLength
	return
}

func EmailIsValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
