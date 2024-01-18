package cmd

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/jonasgao/tomcat-porter/util"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func query(ctx *cli.Context) error {
	xpath := ctx.Args().First()
	path := ctx.Args().Get(1)
	if xpath == "" {
		fmt.Println("There is no XPath parameter.")
		return nil
	}
	var err error
	var file *os.File
	var doc *xmlquery.Node
	if path == "" {
		path, err = util.Search()
		if err != nil {
			return err
		}
	} else if strings.HasSuffix(path, "/") {
		path, err = util.SearchIn(path)
		if err != nil {
			return err
		}
	}
	file, err = os.Open(path)
	if err != nil {
		return err
	}
	doc, err = xmlquery.Parse(file)
	if err != nil {
		return err
	}
	list := xmlquery.Find(doc, xpath)
	for _, node := range list {
		switch node.Type {
		case xmlquery.AttributeNode:
			fmt.Println(node.InnerText())
			break
		case xmlquery.TextNode:
			fmt.Println(node.InnerText())
			break
		default:
			fmt.Printf("Unsupported node type: %d\n", node.Type)
		}
	}
	return nil
}

var Get = cli.Command{
	Name:      "get",
	Aliases:   []string{"g"},
	Usage:     "Get value by XPath with server.xml.",
	UsageText: "tomcat-porter get <xpath> [server.xml file or dir path]",
	Action:    query,
}
