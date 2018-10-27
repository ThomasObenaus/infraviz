package graphml

import "strconv"

var gID int

type Graph struct {
	ID          string      `xml:"id,attr"`
	EdgeDefault EdgeDefault `xml:"edgedefault,attr"`
	Data        []Data      `xml:"data,omitempty"`
	Nodes       []Node      `xml:"node,omitempty"`
}

type EdgeDefault string

const (
	Directed EdgeDefault = "directed"
)

func (g *Graph) AddNode(node Node) {
	g.Nodes = append(g.Nodes, node)
}

func NewGraph() Graph {
	gIDStr := "G" + strconv.Itoa(gID)
	g := Graph{
		EdgeDefault: Directed,
		ID:          gIDStr,
	}
	gID++
	return g
}
