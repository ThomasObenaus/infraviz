package drawyed

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTMLLabel(t *testing.T) {

	label := NewHTMLLabel()
	assert.NotNil(t, label)

	rendered := label.String()
	assert.Equal(t, "<html></html>", rendered)
}

func TestNewLine(t *testing.T) {

	label := NewHTMLLabel()
	assert.NotNil(t, label)
	label.NewLine()

	rendered := label.String()
	assert.Equal(t, "<html>\n  <br></br>\n</html>", rendered)

	label.NewLine()
	rendered = label.String()
	assert.Equal(t, "<html>\n  <br></br>\n  <br></br>\n</html>", rendered)
}

func TestAddText(t *testing.T) {

	label := NewHTMLLabel()
	assert.NotNil(t, label)
	label.AddText("Hallo")

	rendered := label.String()
	assert.Equal(t, "<html>\n  <span>Hallo</span>\n</html>", rendered)

	label.AddText("Welt")

	rendered = label.String()
	assert.Equal(t, "<html>\n  <span>Hallo</span>\n  <span>Welt</span>\n</html>", rendered)
}

func TestAddStyledText(t *testing.T) {

	label := NewHTMLLabel()
	assert.NotNil(t, label)
	label.AddStyledText("Hallo", FSBold)

	rendered := label.String()
	assert.Equal(t, "<html>\n  <b>Hallo</b>\n</html>", rendered)

	label.AddStyledText("Welt", FSItalic)

	rendered = label.String()
	assert.Equal(t, "<html>\n  <b>Hallo</b>\n  <i>Welt</i>\n</html>", rendered)
}

func TestMixed(t *testing.T) {

	label := NewHTMLLabel()
	assert.NotNil(t, label)
	label.AddStyledText("Hallo Welt", FSBold)
	label.NewLine()
	label.AddText("Hello World")
	label.NewLine()

	rendered := label.String()
	assert.Equal(t, "<html>\n  <b>Hallo Welt</b>\n  <br></br>\n  <span>Hello World</span>\n  <br></br>\n</html>", rendered)
}

func ExampleHTMLLabel_AddText() {
	label := NewHTMLLabel()
	label.AddText("Hallo Welt")
	rendered := label.String()
	fmt.Println(rendered)

	// Output: <html>
	//   <span>Hallo Welt</span>
	// </html>
}

func ExampleHTMLLabel_mixed() {
	label := NewHTMLLabel()
	label.AddText("Hallo Welt")
	label.NewLine()
	label.AddStyledText("Hello World", FSBold)
	label.AddStyledText("Hello World", FSItalic)

	rendered := label.String()
	fmt.Println(rendered)

	// Output: <html>
	//   <span>Hallo Welt</span>
	//   <br></br>
	//   <b>Hello World</b>
	//   <i>Hello World</i>
	// </html>
}
