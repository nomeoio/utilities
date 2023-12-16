package utilities

import (
	"bytes"
	"encoding/json"
)

func PrettyJsonString(body []byte) (respJson string, err error) {
	dst := &bytes.Buffer{}
	if err = json.Indent(dst, body, "", "  "); err != nil {
		return
	}
	respJson = dst.String()
	return
}
