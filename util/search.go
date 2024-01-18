package util

import (
	"os"
	"path/filepath"
)

func Search() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	s, err := join(path, "server.xml")
	if err == nil {
		return s, nil
	}
	s, err = join(path, filepath.Join("conf", "server.xml"))
	if err == nil {
		return s, nil
	}
	return "", nil
}

func join(parent string, target string) (string, error) {
	s := filepath.Join(parent, target)
	_, err := os.Stat(s)
	if err == nil {
		return s, nil
	}
	return "", err
}
