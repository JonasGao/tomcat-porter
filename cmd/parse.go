package cmd

import (
	"fmt"
	"github.com/jonasgao/tomcat-porter/parser"
	"github.com/jonasgao/tomcat-porter/util"
	"github.com/urfave/cli/v2"
	"strings"
)

const mode = "mode"
const quite = "quite"

func parse(ctx *cli.Context) error {
	var path string
	var err error
	var servers []parser.Server
	path = ctx.Args().First()
	if path == "" {
		path, err = util.Search()
		if err != nil {
			fmt.Println("Failed search conf/server.xml.")
			return err
		}
		if path == "" {
			fmt.Println("There is no server.xml or conf/server.xml.")
			return nil
		}
		servers, err = parser.One(path)
	} else if strings.HasSuffix(path, "/") {
		servers, err = parser.Dir(path, ctx.Bool(quite))
	} else {
		servers, err = parser.One(path)
	}
	if err != nil {
		return err
	}
	util.Render(servers, ctx.String(mode))
	return nil
}

var Parse = cli.Command{
	Name:    "parse",
	Aliases: []string{"p"},
	Usage:   "Parse server.xml and print all ports.",
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
	Action: parse,
}
