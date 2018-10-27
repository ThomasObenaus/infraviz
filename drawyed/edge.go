package drawyed

type Edge struct {
	ID     string `xml:"id,attr"`
	Source string `xml:"source,attr"`
	Target string `xml:"target,attr"`
	Data   []Data `xml:"data,omitempty"`
}

type PolyLineEdge struct {
	Path      Path      `xml:"y:Path,omitempty"`
	LineStyle LineStyle `xml:"y:LineStyle,omitempty"`
	Arrows    Arrows    `xml:"y:Arrows,omitempty"`
	BendStyle BendStyle `xml:"y:BendStyle,omitempty"`
	EdgeLabel EdgeLabel `xml:"y:EdgeLabel,omitempty"`
}

type Path struct {
	SX float32 `xml:"sx,attr"`
	SY float32 `xml:"sy,attr"`
	TX float32 `xml:"tx,attr"`
	TY float32 `xml:"ty,attr"`
}

type LineStyle struct {
	Color string  `xml:"color,attr"`
	Type  string  `xml:"type,attr"`
	Width float32 `xml:"width,attr"`
}

type Arrows struct {
	Source string `xml:"source,attr"`
	Target string `xml:"target,attr"`
}

type BendStyle struct {
	Smoothed bool `xml:"smoothed,attr"`
}
