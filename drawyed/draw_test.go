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

func TestEdge(t *testing.T) {

	var buf bytes.Buffer
	yedd := NewYedDraw(&buf)
	assert.NotNil(t, yedd)

	r1 := yedd.Rectangle(0, 0, 120, 80, "Rectangle 1")
	assert.NotNil(t, r1)
	r2 := yedd.Rectangle(200, 300, 120, 80, "Rectangle 2")
	assert.NotNil(t, r2)

	label := "Connects r1 and r2"
	e1 := yedd.Edge(r1, r2, label)
	assert.NotNil(t, e1)

	err := yedd.Render()
	assert.NoError(t, err)

	rendered := buf.String()
	assert.Contains(t, rendered, "<edge id=\""+e1.ID+"\" source=\""+r1.ID+"\" target=\""+r2.ID+"\">")
	assert.Contains(t, rendered, label)
}
