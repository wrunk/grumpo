package main

import (
	"fmt"
	"strings"
)

type Page struct {
	BaseDir  string // Usually "pages", used to mkdir -p (mkdirall)
	LinkDir  string // Just BaseDir without the pages for generating links
	FileName string // Actual file name like "index.html"
	Name     string // FileName sans extension so "index"
	Ext      string // ".md" or ".html"
	FullPath string // "pages/index.html"

	BuildDir      string // Sub "pages" for "build", used for writing to
	BuildFullPath string // build/index.html
}

func isValidExt(ext string) bool {
	for _, e := range validExts {
		if e == ext {
			return true
		}
	}
	return false
}

func hasDuplicateName(idx int, path string) bool {
	for i, page := range pages {
		if i == idx { // Don't compare page to itself
			continue
		}
		if page.BaseDir+page.Name == path {
			return true
		}
	}
	return false
}

// Returns empty string on success, error message on failure
func validateIndex() string {

	foundIdx := false
	for i, page := range pages {

		// TODO also check to make sure each sub dir has an index

		if page.BaseDir == "pages" && page.Name == "index" {
			foundIdx = true
		}

		// Only support html and md files (more checking on that later in gen)
		if !isValidExt(page.Ext) {
			return fmt.Sprintf("Found invalid file extension (%s) for page (%#v)", page.Ext, page)
		}
		// Can't have a drafts dir inside pages because we serve
		// drafts directly like http://localhost:9876/drafts/hello/
		if strings.HasPrefix(page.BaseDir, "pages/drafts") ||
			(page.BaseDir+"/"+page.Name) == "pages/drafts" {
			return fmt.Sprintf("You cannot have a file or dir named drafts inside pages")
		}
		// static dir is reserved for other static assets like css, js, images
		if strings.HasPrefix(page.BaseDir, "pages/static") ||
			(page.BaseDir+"/"+page.Name) == "pages/static" {
			return fmt.Sprintf("You cannot have a file or dir named static inside pages")
		}

		// Make sure no duplicated names. This could happen between dir, html md:
		// pages/warren/
		// pages/warren.html
		// pages/warren.md
		// Any two of those would be considered a dup
		if hasDuplicateName(i, page.BaseDir+page.Name) {
			return fmt.Sprintf("Found duplicate page name (%s)", page.BaseDir+page.Name)
		}
	}
	if !foundIdx {
		return fmt.Sprintf("Didn't find an index.[html|md] in pages/")
	}
	return ""
}
