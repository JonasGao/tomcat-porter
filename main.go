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
		version()
		return
	}
	parse(arg1)
}
