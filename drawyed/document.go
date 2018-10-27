package drawyed

import (
	"encoding/xml"
	"io"
	"strconv"
)

// Document represents the toplevel structure for a Graphml (yed) document
type Document struct {
	XMLName      xml.Name `xml:"graphml"`
	Xmlns        string   `xml:"xmlns,attr"`
	XmlnsJava    string   `xml:"xmlns:java,attr"`
	XmlnsSys     string   `xml:"xmlns:sys,attr"`
	XmlnsX       string   `xml:"xmlns:x,attr"`
	XmlnsXsi     string   `xml:"xmlns:xsi,attr"`
	XmlnsY       string   `xml:"xmlns:y,attr"`
	XmlnsYed     string   `xml:"xmlns:yed,attr"`
	XsiSchemaLoc string   `xml:"xsi:schemaLocation,attr"`
	Graphs       []Graph  `xml:"graph"`
	Keys         []Key    `xml:"key,omitempty"`
	meta         meta
	style        style
}

type meta struct {
	keyIDNodeGraphics string // id of the key for yed nodegraphics
	keyIDEdgeGraphics string // id of the key for yed edgegraphics
	lastNodeID        int    // the last used node id
	lastEdgeID        int    // the last used edge id
}

type style struct {
	nodeLabelStyle nodeLabelStyle
	edgeLabelStyle edgeLabelStyle
}

// AddNode adds the given node to the main graph
func (d *Document) AddNode(node Node) {
	d.Graphs[0].AddNode(node)
}

// AddEdge adds the given edge to the main graph
func (d *Document) AddEdge(edge Edge) {
	d.Graphs[0].AddEdge(edge)
}

// AddKey adds a key to the Document
func (d *Document) AddKey(key Key) {
	d.Keys = append(d.Keys, key)
}

// Encode encodes the document to the given writer
func (d *Document) Encode(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")

	w.Write([]byte(xml.Header))

	err := enc.Encode(d)
	if err != nil {
		return err
	}

	return nil
}

// KeyIDNodeGraphics returns the keyID for nodegraphics
func (d *Document) KeyIDNodeGraphics() string {
	return d.meta.keyIDNodeGraphics
}

// KeyIDEdgeGraphics returns the keyID for edgegraphics
func (d *Document) KeyIDEdgeGraphics() string {
	return d.meta.keyIDEdgeGraphics
}

func (d *Document) newGraphID() string {
	return "g" + strconv.Itoa(len(d.Graphs))
}

func (d *Document) newNodeID() string {
	nodeID := "n" + strconv.Itoa(d.meta.lastNodeID)
	d.meta.lastNodeID++
	return nodeID
}

func (d *Document) newKeyID() string {
	return "k" + strconv.Itoa(len(d.Keys))
}

func (d *Document) newEdgeID() string {
	edgeID := "e" + strconv.Itoa(d.meta.lastEdgeID)
	d.meta.lastEdgeID++
	return edgeID
}

// NewEmptyDocument creates an empty Document
func NewEmptyDocument() Document {
	return Document{
		Xmlns:        "http://graphml.graphdrawing.org/xmlns",
		XmlnsJava:    "http://www.yworks.com/xml/yfiles-common/1.0/java",
		XmlnsSys:     "http://www.yworks.com/xml/yfiles-common/markup/primitives/2.0",
		XmlnsX:       "http://www.yworks.com/xml/yfiles-common/markup/2.0",
		XmlnsXsi:     "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsY:       "http://www.yworks.com/xml/graphml",
		XmlnsYed:     "http://www.yworks.com/xml/yed/3",
		XsiSchemaLoc: "http://graphml.graphdrawing.org/xmlns http://www.yworks.com/xml/schema/graphml/1.1/ygraphml.xsd",
	}
}

// NewInitializedDocument creates an initialized Document, having one graph
func NewInitializedDocument() Document {
	doc := NewEmptyDocument()

	doc.style.nodeLabelStyle = defaultNodeLabelStyle
	doc.style.edgeLabelStyle = defaultEdgeLabelStyle

	graph := NewGraph(doc.newGraphID())
	doc.Graphs = append(doc.Graphs, graph)

	// add the key needed for yed graphics
	doc.meta.keyIDNodeGraphics = doc.newKeyID()
	keyNodeGraphics := Key{
		For:    "node",
		YFType: "nodegraphics",
		ID:     doc.meta.keyIDNodeGraphics,
	}
	doc.Keys = append(doc.Keys, keyNodeGraphics)

	// add the key needed for yed edges
	doc.meta.keyIDEdgeGraphics = doc.newKeyID()
	keyEdgeGraphics := Key{
		For:    "edge",
		YFType: "edgegraphics",
		ID:     doc.meta.keyIDEdgeGraphics,
	}
	doc.Keys = append(doc.Keys, keyEdgeGraphics)

	return doc
}
