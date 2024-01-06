package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

const mode = "mode"
const quite = "quite"

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		version()
	}

	app := &cli.App{
		Name:                 "tomcat-porter",
		Usage:                "Parse tomcat server.xml, and print all ports.",
		Version:              Version,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    quite,
				Aliases: []string{"q"},
				Usage:   "Ignore filesystem error.",
			},
			&cli.StringFlag{
				Name:    mode,
				Value:   "list",
				Aliases: []string{"m"},
				Usage:   "Print ports in list/simple/table mode.",
			},
		},
		Action: func(ctx *cli.Context) error {
			var path string
			var err error
			path = ctx.Args().First()
			if path == "" {
				path, err = search()
				if err != nil {
					fmt.Println("Failed search conf/server.xml.")
					return err
				}
				if path == "" {
					fmt.Println("There is no conf/server.xml.")
					return nil
				}
				err := parseOne(path, ctx)
				if err != nil {
					return err
				}
			}
			if strings.HasSuffix(path, "/") {
				parseDir(path, ctx.Bool(quite), ctx)
			} else {
				err := parseOne(path, ctx)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
