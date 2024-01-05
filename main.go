package main

import (
	"fmt"
	"os"
	"strings"
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
		if path == "" {
			fmt.Println("There is no conf/server.xml.")
			return
		}
		parse(path)
		break
	case 2:
		var arg1 = os.Args[1]
		if arg1 == "version" {
			version()
		} else if strings.HasSuffix(arg1, "/") {
			parseDir(arg1, "")
		} else {
			parse(arg1)
		}
		break
	case 3:
		var arg1 = os.Args[1]
		var arg2 = os.Args[2]
		if strings.HasSuffix(arg1, "/") {
			parseDir(arg1, arg2)
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
