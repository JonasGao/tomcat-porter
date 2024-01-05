package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func load(path string) (Server, error) {
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
	return server, err
}

func parse(path string) {
	server, err := load(path)
	if err != nil {
		fmt.Println("Failed load xml file.")
		fmt.Println(err)
		os.Exit(2)
		return
	}
	fmt.Println("Parsing " + path)
	fmt.Println("Server port   : " + server.Port)
	for _, service := range server.Services {
		fmt.Println("  Service " + service.Name)
		for _, connector := range service.Connectors {
			fmt.Println("    Connector : port = " + connector.Port + ", redirectPort = " + connector.RedirectPort)
		}
	}
}

func parseDir(path string, option string) {
	err := filepath.WalkDir(path, visit(option))
	if err != nil {
		fmt.Println("Failed walk dir.")
		fmt.Println(err)
	}
}

func visit(option string) func(string, fs.DirEntry, error) error {
	switch option {
	case "-q":
		return func(path string, d fs.DirEntry, e error) error {
			if e != nil {
				return nil
			}
			if d.Name() == "server.xml" {
				parse(path)
			}
			return nil
		}
	default:
	}
	return func(path string, d fs.DirEntry, e error) error {
		if e != nil {
			fmt.Println("Skip " + path + ", cause: ")
			fmt.Println(e)
			return nil
		}
		if d.Name() == "server.xml" {
			parse(path)
		}
		return nil
	}
}
