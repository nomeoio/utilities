package utilities

import (
	"errors"
	"testing"
)

func TestFunc(t *testing.T) {
	FormatError("Asia/Shanghai", "2006-01-02 15:04:05", errors.New("damn error"))
}
