package graphml

import (
	"encoding/xml"
	"io"
)

func Encode(w io.Writer, doc Document) error {

	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")

	w.Write([]byte(xml.Header))

	err := enc.Encode(doc)
	if err != nil {
		return err
	}

	return nil
}
