package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v2"
	"os"
)

func renderTable(servers []Server) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Server", "Service", "Connector", "Connector"})
	t.AppendHeader(table.Row{"Port", "Name", "Port", "RedirectPort"})
	for _, server := range servers {
		t.AppendSeparator()
		for _, service := range server.Services {
			for _, connector := range service.Connectors {
				t.AppendRow(table.Row{server.Port, service.Name, connector.Port, connector.RedirectPort})
			}
		}
	}
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: true},
	})
	t.SetAutoIndex(true)
	t.Render()
}

func render(server []Server, ctx *cli.Context) {
	switch ctx.String(mode) {
	case "simple":
		renderSimple(server)
		break
	case "table":
		renderTable(server)
		break
	default:
	case "list":
		renderList(server)
	}
}

func renderList(servers []Server) {
	p := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	l := list.NewWriter()
	for _, server := range servers {
		l.AppendItem(server.Path)
		l.Indent()
		l.AppendItem("Server Port: " + p(server.Port))
		l.Indent()
		for _, service := range server.Services {
			l.AppendItem("Service: " + service.Name)
			l.Indent()
			for _, connector := range service.Connectors {
				l.AppendItem("Port: " + p(connector.Port))
				l.AppendItem("RedirectPort: " + p(connector.RedirectPort))
			}
			l.UnIndent()
		}
		l.UnIndent()
		l.UnIndent()
	}
	l.SetOutputMirror(os.Stdout)
	l.SetStyle(list.StyleConnectedRounded)
	l.Render()
}

func renderSimple(servers []Server) {
	for _, server := range servers {
		fmt.Println(server.Port)
		for _, service := range server.Services {
			for _, connector := range service.Connectors {
				fmt.Println(connector.Port, connector.RedirectPort)
			}
		}
	}
}
