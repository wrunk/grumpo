package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// TODO see if we can get a table of contents
func buildMarkdown(mdContent []byte) []byte {

	extensions := parser.NoIntraEmphasis | parser.Tables | parser.FencedCode |
		parser.Autolink | parser.Strikethrough | parser.SpaceHeadings | parser.HeadingIDs |
		parser.BackslashLineBreak | parser.DefinitionLists | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	// TODO confirm this is proper with unicode
	// It should because we're going from disk bytes to md bytes
	return markdown.ToHTML(mdContent, parser, nil)
}
