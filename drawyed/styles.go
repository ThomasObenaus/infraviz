package drawyed

// BorderType represents border types
type BorderType string

const (
	Line BorderType = "line"
)

// ShapeType represents shapes
type ShapeType string

const (
	STypeRectangle ShapeType = "rectangle"
)

// FontStyle represents different font styles
type FontStyle string

const (
	FSBold   FontStyle = "bold"
	FSItalic FontStyle = "italic"
)

// LabelPlacementModel represents different model names for label placement
type LabelPlacementModel string

const (
	LPMInternal LabelPlacementModel = "internal"
)

// LabelModelPosition represents different options for the LabelModels
type LabelModelPosition string

const (
	LMPTop         LabelModelPosition = "t"
	LMPTopLeft     LabelModelPosition = "tl"
	LMPTopRight    LabelModelPosition = "tr"
	LMPBottom      LabelModelPosition = "b"
	LMPBottomLeft  LabelModelPosition = "bl"
	LMPBottomRight LabelModelPosition = "br"
	LMPRight       LabelModelPosition = "r"
	LMPLeft        LabelModelPosition = "l"
	LMPCenter      LabelModelPosition = "c"
)
