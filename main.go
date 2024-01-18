package main

import (
	"github.com/jonasgao/tomcat-porter/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		version()
	}
	app := &cli.App{
		Name:                 "tomcat-porter",
		Usage:                "Parse tomcat server.xml, and print all ports.",
		Version:              Version,
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			&cmd.Get,
			&cmd.Parse,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
