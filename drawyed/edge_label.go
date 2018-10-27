package drawyed

var defaultEdgeLabelStyle = edgeLabelStyle{
	Alignment:  "center",
	FontFamily: "Dialog",
	FontSize:   12,
	FontStyle:  "plain",
}

type edgeLabelStyle struct {
	Alignment  string
	FontFamily string
	FontSize   int
	FontStyle  string
}

func NewEdgeLabel(label string, style edgeLabelStyle) EdgeLabel {
	return EdgeLabel{
		Label:                  label,
		Alignment:              style.Alignment,
		AutoSizePolicy:         "content",
		FontFamily:             style.FontFamily,
		FontSize:               style.FontSize,
		FontStyle:              style.FontStyle,
		HorizontalTextPosition: "center",
		VerticalTextPosition:   "bottom",
	}
}

func NewDefaultEdgeLabel(label string) EdgeLabel {
	return EdgeLabel{
		Label:                  label,
		Alignment:              defaultEdgeLabelStyle.Alignment,
		AutoSizePolicy:         "content",
		FontFamily:             defaultEdgeLabelStyle.FontFamily,
		FontSize:               defaultEdgeLabelStyle.FontSize,
		FontStyle:              defaultEdgeLabelStyle.FontStyle,
		HorizontalTextPosition: "center",
		VerticalTextPosition:   "bottom",
	}
}

type EdgeLabel struct {
	Label                  string `xml:",chardata"`
	Alignment              string `xml:"alignement,attr"`
	AutoSizePolicy         string `xml:"autoSizePolicy,attr"`
	FontFamily             string `xml:"fontFamily,attr"`
	FontSize               int    `xml:"fontSize,attr"`
	FontStyle              string `xml:"fontStyle,attr"`
	HorizontalTextPosition string `xml:"horizontalTextPosition,attr"`
	VerticalTextPosition   string `xml:"verticalTextPosition,attr"`
}
