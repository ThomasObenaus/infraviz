package drawyed

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

func NewGraph(id string) Graph {
	g := Graph{
		EdgeDefault: Directed,
		ID:          id,
	}
	return g
}
