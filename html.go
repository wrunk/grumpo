package main

import (
	"bytes"
	"encoding/xml"
	"io"
	"text/template"
)

// Only do this when NOT testing
func loadBaseTemplate() {
	var err error

	// baseTemplate = template.New("base")
	// Provide a list of custom/helper functions to the template rendering process.
	// See funcs.go
	baseTemplate, err = template.New("base.html").Funcs(templateFuncs).ParseFiles(baseFile)

	if err != nil {
		die("Failed to open base file template %s %s", baseFile, err)
	}
}

// TODO might not be used anymore, should be called execTemplate really
// Will simply render the base template with {{.page}} var
// set to the contents of the current page being rendered or
// displayed localally
func renderHTML(data map[string]interface{}) []byte {

	buf := &bytes.Buffer{}
	err := baseTemplate.Execute(buf, data)
	if err != nil {
		die("Failed to render template: %s", err)
	}
	return buf.Bytes()
}

// TODO find some better lib for this crap srsly
func validateHTML(page []byte) error {
	buf := bytes.NewBuffer(page)
	dec := xml.NewDecoder(buf)

	// Configure the decoder for HTML; leave off strict and autoclose for XHTML
	dec.Strict = false
	dec.AutoClose = xml.HTMLAutoClose
	dec.Entity = xml.HTMLEntity
	for {
		_, err := dec.Token()
		switch err {
		case io.EOF:
			return nil // We're done, it's valid!
		case nil:
		default:
			return err // Oops, something wasn't right
		}
	}
}
