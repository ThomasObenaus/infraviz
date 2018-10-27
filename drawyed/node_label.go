package drawyed

var defaultNodeLabelStyle = nodeLabelStyle{
	Alignment:  "center",
	FontFamily: "Dialog",
	FontSize:   12,
	FontStyle:  "plain",
}

type nodeLabelStyle struct {
	Alignment  string
	FontFamily string
	FontSize   int
	FontStyle  string
}

func NewNodeLabel(label string, style nodeLabelStyle) NodeLabel {
	return NodeLabel{
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

func NewDefaultNodeLabel(label string) NodeLabel {
	return NodeLabel{
		Label:                  label,
		Alignment:              defaultNodeLabelStyle.Alignment,
		AutoSizePolicy:         "content",
		FontFamily:             defaultNodeLabelStyle.FontFamily,
		FontSize:               defaultNodeLabelStyle.FontSize,
		FontStyle:              defaultNodeLabelStyle.FontStyle,
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
