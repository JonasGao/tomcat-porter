package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
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
	val, _ := io.ReadAll(xmlFile)
	var server Server
	err = xml.Unmarshal(val, &server)
	return server, err
}
