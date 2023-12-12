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

func testLestratGoLibxml2(r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")

	p := parser.New(parser.XMLParseNoEnt)

	// BAD: User input used directly in an XPath expression
	_, _ = p.ParseString("//users/user[login/text()='" + username + "']/home_dir/text()")
}
