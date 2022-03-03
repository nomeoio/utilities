package utilities

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
)

func FormatError(timezone, timeFormat string, err error) (errMsg string) {
	var now, wd string
	var e error
	if wd, e = os.Getwd(); e != nil {
		log.Panic(e)
	}
	now, _ = TimeNowString(timezone, timeFormat)
	errMsg = fmt.Sprintf("%s\nerr: %s\ndetail: %s", now, err.Error(), strings.Replace(string(debug.Stack()), wd, ".", -1))
	return
}
