package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// TODO see if we can get a table of contents
func buildMarkdown(mdContent []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	// TODO confirm this is proper with unicode
	// It should because we're going from disk bytes to md bytes
	return markdown.ToHTML(mdContent, parser, nil)
}
