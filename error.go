package utilities

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func FormatError(timezone, timeFormat string, err error) (errMsg string, e error) {
	var now, wd string
	if wd, e = os.Getwd(); e != nil {
		return
	}
	now, _ = TimeNowString(timezone, timeFormat)
	errMsg = fmt.Sprintf("%s\nerr: %s\ndetail: %s", now, err.Error(), strings.Replace(string(debug.Stack()), wd, ".", -1))
	return
}
