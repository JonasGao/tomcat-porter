package main

import (
	"fmt"
	"os"
)

func main() {
	server, err := parse("example.xml")
	if err != nil {
		fmt.Println("Failed parse xml file.")
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Println("Server port is: " + server.Port)
	for _, service := range server.Services {
		for _, connector := range service.Connectors {
			fmt.Println("Connector port is: " + connector.Port)
		}
	}
}
