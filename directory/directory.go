package directory

import (
	"os"
	"path/filepath"
	"strings"
)

func CurrentWorking() string {
	return workingDirectory()
}

func workingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		wd = os.Getenv("PWD")
	}
	wd = filepath.Clean(wd)
	wd = cleanHomeDirectory(wd)

	return wd
}

func cleanHomeDirectory(wd string) string {
	return strings.Replace(wd, os.Getenv("HOME"), "~", 1)
}
