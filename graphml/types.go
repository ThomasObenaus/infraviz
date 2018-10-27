package graphml

import "encoding/xml"

type Document struct {
	XMLName      xml.Name `xml:"graphml"`
	Xmlns        string   `xml:"xmlns,attr"`
	XmlnsJava    string   `xml:"xmlns:java,attr"`
	XmlnsSys     string   `xml:"xmlns:sys,attr"`
	XmlnsX       string   `xml:"xmlns:x,attr"`
	XmlnsXsi     string   `xml:"xmlns:xsi,attr"`
	XmlnsY       string   `xml:"xmlns:y,attr"`
	XmlnsYed     string   `xml:"xmlns:yed,attr"`
	XsiSchemaLoc string   `xml:"xsi:schemaLocation,attr"`
	Graphs       []Graph  `xml:"graph"`
	Keys         []Key    `xml:"key,omitempty"`
}

func NewEmptyDocument() Document {
	return Document{
		Xmlns:        "http://graphml.graphdrawing.org/xmlns",
		XmlnsJava:    "http://www.yworks.com/xml/yfiles-common/1.0/java",
		XmlnsSys:     "http://www.yworks.com/xml/yfiles-common/markup/primitives/2.0",
		XmlnsX:       "http://www.yworks.com/xml/yfiles-common/markup/2.0",
		XmlnsXsi:     "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsY:       "http://www.yworks.com/xml/graphml",
		XmlnsYed:     "http://www.yworks.com/xml/yed/3",
		XsiSchemaLoc: "http://graphml.graphdrawing.org/xmlns http://www.yworks.com/xml/schema/graphml/1.1/ygraphml.xsd",
	}
}

type EdgeDefault string

const (
	Directed EdgeDefault = "directed"
)

type Key struct {
	ID       string `xml:"id,attr,omitempty"`
	AttrName string `xml:"attr.name,attr,omitempty"`
	For      string `xml:"for,attr,omitempty"`
	YFType   string `xml:"yfiles.type,attr,omitempty"`
}

type Graph struct {
	ID          string      `xml:"id,attr"`
	EdgeDefault EdgeDefault `xml:"edgedefault,attr"`
	Data        []Data      `xml:"data,omitempty"`
	Nodes       []Node      `xml:"node,omitempty"`
}

type Data struct {
	Key       string     `xml:"key,attr"`
	ShapeNode *ShapeNode `xml:"y:ShapeNode,omitempty"`
}

type Node struct {
	ID   string `xml:"id,attr"`
	Data []Data `xml:"data,omitempty"`
}
