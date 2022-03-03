package utilities

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func HttpRequest(requestMethod string, reqBody []byte, url string, headers [][]string) (respBody []byte, err error) {
	var req *http.Request
	var resp *http.Response
	req, err = http.NewRequest(requestMethod, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return
	}
	if len(headers) > 0 {
		for _, header := range headers {
			req.Header.Add(header[0], header[1])
		}
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	var buf *bytes.Buffer = new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return
	}

	respBody = buf.Bytes()
	return
}

func DownloadFile(url, fn string, ignoreErr bool) (err error) {

	// Create blank file
	var file *os.File
	if file, err = os.Create(fn); err != nil {
		return
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	// Put content on file
	var resp *http.Response
	if resp, err = client.Get(url); err != nil {
		return
	}
	defer resp.Body.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return
	}
	defer file.Close()
	return
}
