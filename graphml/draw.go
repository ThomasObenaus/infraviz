package graphml

import "io"

type GMLDraw struct {
	doc Document
}

func (d *GMLDraw) Rectangle(x float32, y float32, width float32, height float32) {

}

func (d *GMLDraw) Render(w io.Writer) error {
	if err := d.doc.Encode(w); err != nil {
		return err
	}
	return nil
}

func NewGMLDraw(w io.Writer) GMLDraw {
	return GMLDraw{
		doc: NewEmptyDocument(),
	}
}
