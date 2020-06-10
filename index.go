package main

import (
	"fmt"
	"strings"
	"time"
)

/*
"Index" means our list (slice) of pages
*/

type Page struct {
	BaseDir  string // Usually "pages", used to mkdir -p (mkdirall)
	LinkDir  string // Just BaseDir without the pages for generating links
	FileName string // Actual file name like "index.html"
	Name     string // FileName sans extension so "index"
	Ext      string // ".md" or ".html"
	FullPath string // "pages/index.html"

	BuildDir      string // Sub "pages" for "build", used for writing to
	BuildFullPath string // build/index.html

	Meta *Meta
}

func (p Page) LocalLink() string {
	return fmt.Sprintf("/%s/", p.LinkDir)
}

// Means this page's publish date is in the past
func (p Page) Published() bool {
	return time.Now().After(time.Time(p.Meta.PublishDate))
}

// TODO need to figure out error handling here...
// This is simply providing the page content to get injected into
// the "page" template block like:
/*
{{- block "page" }}
<h1>{{ .page.Meta.Title }}</h1>
{{ .page.HTML }} {{ -}}
*/
func (p Page) HTML() string {

	pageDataBys, err := readOffsetFile(p.FullPath, p.Meta.contentStartsOn)
	if err != nil {
		panic(fmt.Sprintf("Couldn't read file %s %s", p.FullPath, err))
	}
	if p.Ext == extMarkdown {
		pageDataBys = buildMarkdown(pageDataBys)
	}

	// TODO, support OPTIONALLY overriding the page block for things like the home
	// page which will want to use go templates to show recent posts
	if p.Meta.RenderGoTemplate {
		pageDataBys = goRenderPage(p, pageDataBys)
	}

	// Make sure the generated/final html is valid
	err = validateHTML(pageDataBys)
	if err != nil {
		panic(fmt.Sprintf("%s Resulted in invalid html (%s)", p.FullPath, err))
	}
	return string(pageDataBys)
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

// validateIndex checks all pages for valid extensions and
// checks various other things to make sure we didnt screw ourselves
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
