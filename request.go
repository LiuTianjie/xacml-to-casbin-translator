package main

import (
	"encoding/xml"
)

type Request struct {
	XMLName     xml.Name    `xml:"Request"`
	Subject     Subject     `xml:"Subject"`
	Resource    Resource    `xml:"Resource"`
	Action      Action      `xml:"Action"`
	Environment Environment `xml:"Environment"`
}

type Subject struct {
	XMLName   xml.Name    `xml:"Subject"`
	Attribute []Attribute `xml:"Attribute"`
}

type Resource struct {
	XMLName   xml.Name    `xml:"Resource"`
	Attribute []Attribute `xml:"Attribute"`
}

type Action struct {
	XMLName   xml.Name    `xml:"Action"`
	Attribute []Attribute `xml:"Attribute"`
}

type Environment struct {
	XMLName   xml.Name    `xml:"Environment"`
	Attribute []Attribute `xml:"Attribute"`
}

type Attribute struct {
	XMLName        xml.Name `xml:"Attribute" json:"-"`
	AttributeValue string   `xml:"AttributeValue"`
	AttributeId    string   `xml:"AttributeId,attr"`
}
