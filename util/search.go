package util

import (
	"errors"
	"os"
	"path/filepath"
)

func Search() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return SearchIn(path)
}

func joinAndStat(parent string, target string) (string, error) {
	s := filepath.Join(parent, target)
	_, err := os.Stat(s)
	if err == nil {
		return s, nil
	}
	return "", err
}

func SearchIn(path string) (string, error) {
	s, err := joinAndStat(path, "server.xml")
	if err == nil {
		return s, nil
	}
	s, err = joinAndStat(path, filepath.Join("conf", "server.xml"))
	if err == nil {
		return s, nil
	}
	return "", errors.New("there is no server.xml or conf/server.xml")
}

func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}
