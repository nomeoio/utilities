package utilities

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(time.RFC3339)
}

func TestTimeNowString(t *testing.T) {
	var nowString string
	var err error
	if nowString, err = TimeNowString("Asia/Shanghai", time.RFC3339); err != nil {
		return
	}
	t.Log(nowString)
}
