package main

import (
	"encoding/xml"
	"fmt"
	"github.com/urfave/cli/v2"
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

func parseOne(path string, ctx *cli.Context) error {
	server, err := parse(path)
	if err != nil {
		fmt.Println("Failed load xml file: " + path)
		fmt.Println(err)
		return err
	}
	render([]Server{server}, ctx)
	return nil
}

func parseDir(path string, quite bool, cCtx *cli.Context) {
	servers := make([]Server, 0)
	err := filepath.WalkDir(path, visit(&servers, quite))
	if err != nil {
		fmt.Println("Failed walk dir.")
		fmt.Println(err)
	} else {
		render(servers, cCtx)
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
