package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Error: there is no input file path.")
		os.Exit(1)
		return
	}
	var arg1 = os.Args[1]
	if arg1 == "version" {
		fmt.Println("Version: 1")
	}
	server, err := parse(arg1)
	if err != nil {
		fmt.Println("Failed parse xml file.")
		fmt.Println(err)
		os.Exit(2)
		return
	}
	fmt.Print("Using port has, server = " + server.Port)
	for _, service := range server.Services {
		for _, connector := range service.Connectors {
			fmt.Print(", connector = " + connector.Port + ", redirectPort = " + connector.RedirectPort)
		}
	}
	fmt.Println()
}
