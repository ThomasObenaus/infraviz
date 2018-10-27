package drawyed

import "io"

type YedDraw struct {
	doc    Document
	writer io.Writer
}

func (d *YedDraw) Rectangle(x float32, y float32, width float32, height float32, label string) {

	doc := &d.doc

	nodeKeyID := doc.KeyIDNodeGraphics()
	node := Node{ID: doc.newNodeID()}

	snode := &ShapeNode{}
	snode.Geometry = Geometry{Height: height, Width: width, X: x, Y: y}
	snode.Shape = Shape{ShapeType: Rectangle}
	nodeLabel := NewNodeLabel(label, doc.style.nodeLabelStyle)
	snode.NodeLabel = &nodeLabel
	node.Data = []Data{Data{Key: nodeKeyID, ShapeNode: snode}}

	doc.AddNode(node)
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
