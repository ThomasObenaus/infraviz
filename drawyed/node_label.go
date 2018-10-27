package drawyed

var defaultLabelStyle = labelStyle{
	Alignment:  "center",
	FontFamily: "Dialog",
	FontSize:   12,
	FontStyle:  "plain",
}

type labelStyle struct {
	Alignment  string
	FontFamily string
	FontSize   int
	FontStyle  string
}

func NewNodeLabel(label string, style labelStyle) NodeLabel {
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
		Alignment:              defaultLabelStyle.Alignment,
		AutoSizePolicy:         "content",
		FontFamily:             defaultLabelStyle.FontFamily,
		FontSize:               defaultLabelStyle.FontSize,
		FontStyle:              defaultLabelStyle.FontStyle,
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
