package utilities

import (
	"bytes"
	"encoding/json"
	"log"
)

func PrettyJsonString(body []byte) (respJson string) {
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, body, "", "  "); err != nil {
		log.Panic(err)
	}
	respJson = dst.String()
	return
}
