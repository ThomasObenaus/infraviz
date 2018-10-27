package graphml

type Key struct {
	ID       string `xml:"id,attr,omitempty"`
	AttrName string `xml:"attr.name,attr,omitempty"`
	For      string `xml:"for,attr,omitempty"`
	YFType   string `xml:"yfiles.type,attr,omitempty"`
}

type Data struct {
	Key       string     `xml:"key,attr"`
	ShapeNode *ShapeNode `xml:"y:ShapeNode,omitempty"`
}

type Node struct {
	ID   string `xml:"id,attr"`
	Data []Data `xml:"data,omitempty"`
}
