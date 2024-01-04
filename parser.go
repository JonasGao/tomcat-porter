package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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
	b, _ := ioutil.ReadAll(xmlFile)
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
	fmt.Println("Server port   : " + server.Port)
	for _, service := range server.Services {
		fmt.Println("  Service " + service.Name)
		for _, connector := range service.Connectors {
			fmt.Println("    Connector : port = " + connector.Port + ", redirectPort = " + connector.RedirectPort)
		}
	}
	fmt.Println()
}
