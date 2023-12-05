package vulns

import (
	"fmt"
	"net/http"

	"github.com/lestrrat-go/libxml2/parser"
)

func Xxe(w http.ResponseWriter, r *http.Request) {
	xmlData := r.URL.Query().Get("xmlData")

	if xmlData == "" {
		http.Error(w, "Missing xmlData parameter", http.StatusBadRequest)
		return
	}

	p := parser.New(parser.XMLParseNoEnt)
	// xmlData="<!DOCTYPE d [<!ENTITY e SYSTEM \"file:///etc/passwd\">]><t>&e;</t>"
	doc, err := p.ParseString(xmlData)

	defer doc.Free()

	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing XML: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Parsed user: %+v\n", doc)
}
