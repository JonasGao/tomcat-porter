package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func parse(path string) (Server, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return emptyServer(), err
	}
	defer func(xmlFile *os.File) {
		err := xmlFile.Close()
		if err != nil {
			fmt.Println("Failed closed xml file.")
			fmt.Println(err)
		}
	}(xmlFile)
	b, _ := io.ReadAll(xmlFile)
	var server Server
	err = xml.Unmarshal(b, &server)
	server.Path = path
	return server, err
}

func One(path string) ([]Server, error) {
	server, err := parse(path)
	if err != nil {
		fmt.Println("Failed load xml file: " + path)
		fmt.Println(err)
		return nil, err
	}
	return []Server{server}, nil
}

func Dir(path string, quite bool) ([]Server, error) {
	servers := make([]Server, 0)
	err := filepath.WalkDir(path, visit(&servers, quite))
	if err != nil {
		fmt.Println("Failed walk dir.")
		fmt.Println(err)
		return nil, err
	} else {
		return servers, nil
	}
}

func visit(servers *[]Server, quite bool) func(string, fs.DirEntry, error) error {
	if quite {
		return func(path string, d fs.DirEntry, e error) error {
			if e != nil {
				return nil
			}
			if d.Name() == "server.xml" {
				server, err := parse(path)
				if err == nil {
					*servers = append(*servers, server)
				}
			}
			return nil
		}
	}
	return func(path string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println("Skip " + path + ", cause: ")
			fmt.Println(e)
			return nil
		}
		if d.Name() == "server.xml" {
			server, err := parse(path)
			if err == nil {
				*servers = append(*servers, server)
			} else {
				return err
			}
		}
		return nil
	}
}
