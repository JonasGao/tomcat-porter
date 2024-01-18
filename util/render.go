package util

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jonasgao/tomcat-porter/parser"
	"os"
)

func renderTable(servers []parser.Server) {
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

func Render(server []parser.Server, mode string) {
	switch mode {
	case "simple":
		renderSimple(server)
		return
	case "table":
		renderTable(server)
		return
	case "list":
		renderList(server)
		return
	}
}

func renderList(servers []parser.Server) {
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

func renderSimple(servers []parser.Server) {
	for _, server := range servers {
		fmt.Println(server.Port)
		for _, service := range server.Services {
			for _, connector := range service.Connectors {
				fmt.Println(connector.Port, connector.RedirectPort)
			}
		}
	}
}
