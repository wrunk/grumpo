package main

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"
	"time"
)

var templateFuncs = template.FuncMap{
	"latest": fnLatestPages,
	"maxstr": fnMaxString,
	"date":   fnDate,
}

func fnDate(ld localDate) string {
	t := time.Time(ld)
	return t.Format("Jan 02, 2006")
}

// Currently used to hold template rendering functions

// Gives the list of latest pages, newest on top with a limit
// num. Uses latestPages as a cache
func fnLatestPages(num int) []Page {
	if len(latestPages) == 0 {
		// Sort descending by publish date into tmp slice
		tmpLatestPages := make([]Page, len(pages))
		copy(tmpLatestPages, pages)
		sort.Slice(tmpLatestPages, func(i, j int) bool {
			return time.Time(tmpLatestPages[i].Meta.PublishDate).After(time.Time(tmpLatestPages[j].Meta.PublishDate))

		})
		// Then loop over tmp slice to filter out excluded
		latestPages = []Page{}
		for i := range tmpLatestPages {
			p := tmpLatestPages[i] // Get the actual slice element, not a copy
			if !p.Meta.ExcludeFromLatest {
				latestPages = append(latestPages, p)
			}
		}
	}

	if len(latestPages) <= num {
		return latestPages
	}
	return latestPages[:num]
}

// Gives back the string with a max len of m
func fnMaxString(s string, m int) string {
	if len(s) <= m {
		return s
	}
	return s[:m]
}

// Render helper to render an individual page
func goRenderPage(page Page, dataBys []byte) []byte {
	t, err := template.New("fromFile").Funcs(templateFuncs).Parse(string(dataBys))
	if err != nil {
		panic(fmt.Sprintf("Failed to render page template (%s)", err))
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, templateData(page))
	if err != nil {
		panic(fmt.Sprintf("Failed to render template: (%s)", err))
	}
	return buf.Bytes()
}

// Standard data to render all template with
func templateData(page Page) map[string]interface{} {
	return map[string]interface{}{
		"page":  page,
		"pages": pages,
	}
}
