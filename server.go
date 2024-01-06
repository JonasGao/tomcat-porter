package main

import "encoding/xml"

type Server struct {
	XMLName  xml.Name  `xml:"Server"`
	Services []Service `xml:"Service"`
	Port     string    `xml:"port,attr"`
	Path     string
}

type Service struct {
	XMLName    xml.Name    `xml:"Service"`
	Name       string      `xml:"name,attr"`
	Connectors []Connector `xml:"Connector"`
}

type Connector struct {
	XMLName      xml.Name `xml:"Connector"`
	Port         string   `xml:"port,attr"`
	Protocol     string   `xml:"protocol,attr"`
	RedirectPort string   `xml:"redirectPort,attr"`
}

func emptyServer() Server {
	return Server{
		XMLName:  xml.Name{},
		Services: nil,
		Port:     "Unknown",
	}
}
