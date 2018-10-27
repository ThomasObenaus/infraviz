package drawyed

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRectangle(t *testing.T) {

	var buf bytes.Buffer
	yedd := NewYedDraw(&buf)
	assert.NotNil(t, yedd)

	yedd.Rectangle(50, 20, 120, 80, "SmallLabel")
	err := yedd.Render()
	assert.NoError(t, err)

	rendered := buf.String()
	assert.Regexp(t, ".*<y:Geometry height=\"80\" width=\"120\" x=\"50\" y=\"20\"></y:Geometry>.*", rendered)
	assert.Contains(t, rendered, "SmallLabel")
}
