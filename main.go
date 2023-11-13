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
	fmt.Print("Using port has, server = " + server.Port)
	for _, service := range server.Services {
		for _, connector := range service.Connectors {
			fmt.Print(", connector = " + connector.Port)
		}
	}
	fmt.Println()
}
