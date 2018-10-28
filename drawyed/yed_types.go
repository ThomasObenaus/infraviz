package drawyed

type ShapeNode struct {
	Geometry    Geometry     `xml:"y:Geometry,omitempty"`
	Fill        *Fill        `xml:"y:Fill,omitempty"`
	BorderStyle *BorderStyle `xml:"y:BorderStyle,omitempty"`
	Shape       Shape        `xml:"y:Shape,omitempty"`
	NodeLabel   *NodeLabel   `xml:"y:NodeLabel,omitempty"`
}

type Geometry struct {
	Height float32 `xml:"height,attr"`
	Width  float32 `xml:"width,attr"`
	X      float32 `xml:"x,attr"`
	Y      float32 `xml:"y,attr"`
}

type Fill struct {
	Color       string `xml:"color,attr"`
	Transparent bool   `xml:"transparent,attr"`
}

type BorderStyle struct {
	Color  string     `xml:"color,attr"`
	Raised bool       `xml:"raised,attr"`
	Type   BorderType `xml:"type,attr"`
	Width  float32    `xml:"width,attr"`
}

type Shape struct {
	ShapeType ShapeType `xml:"type,attr"`
}
