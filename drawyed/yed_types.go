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

type BorderType string

const (
	Line BorderType = "line"
)

type BorderStyle struct {
	Color  string     `xml:"color,attr"`
	Raised bool       `xml:"raised,attr"`
	Type   BorderType `xml:"type,attr"`
	Width  float32    `xml:"width,attr"`
}

type ShapeType string

const (
	Rectangle ShapeType = "rectangle"
)

type Shape struct {
	ShapeType ShapeType `xml:"type,attr"`
}

func NewNodeLabel(label string) NodeLabel {
	return NodeLabel{
		Label:                  label,
		Alignment:              "center",
		AutoSizePolicy:         "content",
		FontFamily:             "Dialog",
		FontSize:               12,
		FontStyle:              "plain",
		HorizontalTextPosition: "center",
		VerticalTextPosition:   "bottom",
	}
}

type NodeLabel struct {
	Label                  string `xml:",chardata"`
	Alignment              string `xml:"alignement,attr"`
	AutoSizePolicy         string `xml:"autoSizePolicy,attr"`
	FontFamily             string `xml:"fontFamily,attr"`
	FontSize               int    `xml:"fontSize,attr"`
	FontStyle              string `xml:"fontStyle,attr"`
	HorizontalTextPosition string `xml:"horizontalTextPosition,attr"`
	VerticalTextPosition   string `xml:"verticalTextPosition,attr"`
}

func NewYedNodeKey(id string) Key {
	return Key{ID: id, For: "node", YFType: "nodegraphics"}
}
