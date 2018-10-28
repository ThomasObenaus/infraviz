package drawyed

var defaultNodeLabelStyle = nodeLabelStyle{
	Alignment:           "center",
	FontFamily:          "Dialog",
	FontSize:            12,
	FontStyle:           "plain",
	LabelPlacementModel: LPMInternal,
	LabelModelPosition:  LMPCenter,
}

type nodeLabelStyle struct {
	Alignment           string
	FontFamily          string
	FontSize            int
	FontStyle           string
	LabelPlacementModel LabelPlacementModel
	LabelModelPosition  LabelModelPosition
}

func NewNodeLabel(label string, style nodeLabelStyle) NodeLabel {
	nodeLabel := NewDefaultNodeLabel(label)
	nodeLabel.Alignment = style.Alignment
	nodeLabel.FontFamily = style.FontFamily
	nodeLabel.FontSize = style.FontSize
	nodeLabel.FontStyle = style.FontStyle
	nodeLabel.LabelPlacementModel = style.LabelPlacementModel
	nodeLabel.LabelModelPosition = style.LabelModelPosition
	return nodeLabel
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
		LabelPlacementModel:    defaultNodeLabelStyle.LabelPlacementModel,
		LabelModelPosition:     defaultNodeLabelStyle.LabelModelPosition,
	}
}

type NodeLabel struct {
	Label                  string              `xml:",chardata"`
	Alignment              string              `xml:"alignement,attr"`
	AutoSizePolicy         string              `xml:"autoSizePolicy,attr"`
	FontFamily             string              `xml:"fontFamily,attr"`
	FontSize               int                 `xml:"fontSize,attr"`
	FontStyle              string              `xml:"fontStyle,attr"`
	HorizontalTextPosition string              `xml:"horizontalTextPosition,attr"`
	VerticalTextPosition   string              `xml:"verticalTextPosition,attr"`
	LabelPlacementModel    LabelPlacementModel `xml:"modelName,attr"`
	LabelModelPosition     LabelModelPosition  `xml:"modelPosition,attr"`
}
