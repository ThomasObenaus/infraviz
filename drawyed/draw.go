package drawyed

import (
	"io"
)

type YedDraw struct {
	doc    Document
	writer io.Writer
}

type Primitive struct {
	ID string
}

func (d *YedDraw) Rectangle(x float32, y float32, width float32, height float32, label string) Primitive {

	doc := &d.doc

	nodeKeyID := doc.KeyIDNodeGraphics()
	nodeID := doc.newNodeID()
	node := Node{ID: nodeID}

	snode := &ShapeNode{}
	snode.Geometry = Geometry{Height: height, Width: width, X: x, Y: y}
	snode.Shape = Shape{ShapeType: STypeRectangle}
	nodeLabel := NewNodeLabel(label, *doc.CurrentNodeLabelStyle())
	snode.NodeLabel = &nodeLabel
	node.Data = []Data{Data{Key: nodeKeyID, ShapeNode: snode}}

	doc.AddNode(node)

	return Primitive{
		ID: nodeID,
	}
}

func (d *YedDraw) Edge(p1 Primitive, p2 Primitive, label string) Primitive {
	doc := &d.doc

	edgeKeyID := doc.KeyIDEdgeGraphics()
	edgeID := doc.newEdgeID()
	edge := Edge{ID: edgeID, Source: p1.ID, Target: p2.ID}

	path := Path{SX: 0, SY: 0, TX: 0, TY: 0}
	lineStyle := LineStyle{Color: "#000000", Type: "line", Width: 1}
	arrows := Arrows{Source: "none", Target: "none"}
	bendStyle := BendStyle{Smoothed: false}
	edgeLabel := NewEdgeLabel(label, doc.style.edgeLabelStyle)

	pEdge := PolyLineEdge{Path: path, LineStyle: lineStyle, Arrows: arrows, BendStyle: bendStyle, EdgeLabel: edgeLabel}
	edge.Data = append(edge.Data, Data{Key: edgeKeyID, PolyLineEdge: &pEdge})

	doc.AddEdge(edge)
	return Primitive{
		ID: edgeID,
	}
}

func (d *YedDraw) Render() error {
	if err := d.doc.Encode(d.writer); err != nil {
		return err
	}
	return nil
}

func NewYedDraw(w io.Writer) YedDraw {
	return YedDraw{
		doc:    NewInitializedDocument(),
		writer: w,
	}
}
