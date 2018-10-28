package drawyed

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInitializedDocument(t *testing.T) {

	doc := NewInitializedDocument()
	assert.Equal(t, 1, len(doc.Graphs))
	assert.Equal(t, 2, len(doc.Keys))
}

func TestAddNode(t *testing.T) {

	doc := NewInitializedDocument()
	node := Node{}
	doc.AddNode(node)
	doc.AddNode(node)
	assert.Equal(t, 2, len(doc.Graphs[0].Nodes))
}

func TestAddKey(t *testing.T) {

	doc := NewInitializedDocument()
	key := Key{}
	doc.AddKey(key)
	doc.AddKey(key)
	assert.Equal(t, 4, len(doc.Keys))
}

func TestEncode(t *testing.T) {
	doc := NewInitializedDocument()

	var buf bytes.Buffer
	err := doc.Encode(&buf)
	assert.NoError(t, err)

	rendered := buf.String()
	assert.Contains(t, rendered, xml.Header)
	assert.Contains(t, rendered, "http://graphml.graphdrawing.org/xmlns")
	assert.Contains(t, rendered, "<graph id=\"g0\" edgedefault=\"directed\">")
}

func TestNewNodeID(t *testing.T) {

	doc := NewEmptyDocument()
	nID := doc.newNodeID()
	assert.Equal(t, "n0", nID)

	nID = doc.newNodeID()
	assert.Equal(t, "n1", nID)
}

func TestNodeLabelStyleChange(t *testing.T) {

	doc := NewEmptyDocument()

	nStyle := doc.PopNodeLabelStyle()
	assert.Nil(t, nStyle)

	nStyle = &nodeLabelStyle{Alignment: "ABC"}
	doc.PushNodeLabelStyle(nStyle)

	nStyle = doc.CurrentNodeLabelStyle()
	assert.NotNil(t, nStyle)
	assert.Equal(t, "ABC", nStyle.Alignment)
}
