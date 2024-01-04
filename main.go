package main

import (
	"fmt"
	"os"
)

func main() {
	switch len(os.Args) {
	case 1:
		path, err := search()
		if err != nil {
			fmt.Println("Failed search conf/server.xml.")
			fmt.Println(err)
			return
		}
		parse(path)
		break
	case 2:
		var arg1 = os.Args[1]
		if arg1 == "version" {
			version()
		} else {
			parse(arg1)
		}
		break
	default:
		fmt.Println("Error: Wrong args size. Just one or no arg.")
		fmt.Println("  example \"tomcat-porter conf/server.xml\"")
		os.Exit(1)
	}
}
