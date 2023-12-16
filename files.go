package utilities

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func AbsCwd() (cwd string, err error) {
	// os.Getwd() reuturns where you're in terminal window.
	// this func returns the directory of the executable
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		err = errors.New("unable to get the current filename")
		return
	}
	cwd = filepath.Dir(filename)
	return
}

func ReadFile(fp string) (b []byte, err error) {
	b, err = os.ReadFile(fp)
	return
}

func WriteFile(b []byte, filename string) (err error) {
	return os.WriteFile(filename, b, 0644)
}

func DoesFileExist(fp string) (exist bool, err error) {
	if _, err = os.Stat(fp); err == nil {
		exist = true
	} else if errors.Is(err, os.ErrNotExist) {
		exist = false
	}
	return
}

func CreateDirIfNotExist(fp string) (err error) {
	var dp string = filepath.Dir(fp) // dir path
	if _, err = os.Stat(dp); os.IsNotExist(err) {
		if err != nil {
			return
		}
		err = os.MkdirAll(dp, os.ModePerm)
	}
	return
}
