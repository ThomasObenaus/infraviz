package drawyed

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyDocument(t *testing.T) {

	doc := NewEmptyDocument()
	assert.NotNil(t, doc)
}

func TestNewInitializedDocument(t *testing.T) {

	doc := NewInitializedDocument()
	assert.NotNil(t, doc)
	assert.Equal(t, 1, len(doc.Graphs))
	assert.Equal(t, 1, len(doc.Keys))
}

func TestAddNode(t *testing.T) {

	doc := NewInitializedDocument()
	assert.NotNil(t, doc)
	node := Node{}
	doc.AddNode(node)
	doc.AddNode(node)
	assert.Equal(t, 2, len(doc.Graphs[0].Nodes))
}

func TestAddKey(t *testing.T) {

	doc := NewInitializedDocument()
	assert.NotNil(t, doc)
	key := Key{}
	doc.AddKey(key)
	doc.AddKey(key)
	assert.Equal(t, 3, len(doc.Keys))
}

func TestEncode(t *testing.T) {
	doc := NewInitializedDocument()
	assert.NotNil(t, doc)

	var buf bytes.Buffer
	err := doc.Encode(&buf)
	assert.NoError(t, err)

	rendered := buf.String()
	assert.Contains(t, rendered, xml.Header)
	assert.Contains(t, rendered, "http://graphml.graphdrawing.org/xmlns")
	assert.Contains(t, rendered, "<graph id=\"g0\" edgedefault=\"directed\">")
}
