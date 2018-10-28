package drawyed

import (
	"bytes"
	"encoding/xml"
)

// HTMLLabel is an interface to generate styled text lables based on html
type HTMLLabel interface {
	// NewLine adds <br></br>
	NewLine()

	// AddStyledText adds <b>text</b> or <i>text</i>
	AddStyledText(text string, fontStyle FontStyle)

	// AddText adds plain text
	AddText(text string)

	// String renders the label to a string
	String() string
}

type htmlLabelImpl struct {
	XMLName xml.Name  `xml:"html"`
	Element []element `xml:",any,omitempty"`
}

type element struct {
	XMLName xml.Name
	Text    string `xml:",chardata"`
}

func (hl *htmlLabelImpl) NewLine() {
	e := element{XMLName: xml.Name{Local: "br"}}
	hl.Element = append(hl.Element, e)
}

func (hl *htmlLabelImpl) AddStyledText(text string, fontStyle FontStyle) {
	e := element{Text: text}

	switch fontStyle {
	case FSBold:
		e.XMLName.Local = "b"
	case FSItalic:
		e.XMLName.Local = "i"
	}
	hl.Element = append(hl.Element, e)
}

func (hl *htmlLabelImpl) AddText(text string) {
	e := element{Text: text}
	e.XMLName.Local = "span"
	hl.Element = append(hl.Element, e)
}

func (hl *htmlLabelImpl) String() string {

	var buf bytes.Buffer
	enc := xml.NewEncoder(&buf)
	enc.Indent("", "  ")

	err := enc.Encode(hl)
	if err != nil {
		return ""
	}

	return buf.String()
}

// NewHTMLLabel creates a new empty html label
func NewHTMLLabel() HTMLLabel {
	return &htmlLabelImpl{}
}
