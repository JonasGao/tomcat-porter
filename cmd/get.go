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
		// Auto search server.xml
		path, err = util.Search()
		if err != nil {
			return err
		}
	} else if strings.HasSuffix(path, "/") {
		// The path is dir explicitly , and search server.xml in dir
		path, err = util.SearchIn(path)
		if err != nil {
			return err
		}
	} else {
		// Check the path
		isDir, err := util.IsDir(path)
		if err != nil {
			return err
		}
		if isDir {
			path, err = util.SearchIn(path)
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
			fmt.Printf("Unsupported node '%s', type is '%s', prefix is '%s', ns is '%s'\n",
				node.Data, nameOfNodeType(node.Type), node.Prefix, node.NamespaceURI)
		}
	}
	return nil
}

func nameOfNodeType(t xmlquery.NodeType) string {
	switch t {
	case xmlquery.TextNode:
		return "text"
	case xmlquery.DocumentNode:
		return "document"
	case xmlquery.DeclarationNode:
		return "declaration"
	case xmlquery.ElementNode:
		return "element"
	case xmlquery.CharDataNode:
		return "char data"
	case xmlquery.CommentNode:
		return "comment"
	case xmlquery.AttributeNode:
		return "attribute"
	default:
		return "Unknown Type"
	}
}

var Get = cli.Command{
	Name:      "get",
	Aliases:   []string{"g"},
	Usage:     "Get value by XPath with server.xml.",
	UsageText: "tomcat-porter get <xpath> [server.xml file or dir path]",
	Action:    query,
}
