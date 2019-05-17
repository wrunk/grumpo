package main

import (
	"io/ioutil"
	"os"
)

// Even if you delete a lot of this stuff, this file
// could be interesting to read through to see if you have
// all standard files for a publisher.

var (
	initialFiles = map[string]string{
		".gitignore":          gitIgnore,
		"robots":              robots,
		"rss":                 rss,
		"sitemap.xml":         sitemapXML,
		"base.html":           baseHTML,
		"static/css/main.css": mainCSS,

		"pages/index.md":       pagesIndexMD,
		"pages/demo/index.md":  pagesDemoIndexMD,
		"pages/demo/help.html": pagesDemoHelpHTML,

		"drafts/index.html":      draftsIndexHTML,
		"drafts/demo/index.md":   draftsDemoIndexMD,
		"drafts/demo/about.html": draftsDemoAboutHTML,
	}
)

// TODO leverage this in other spots that use os.MkdirAll
// Make this file's base dir (no error if already exists,
// same behavior as mkdir -p)
func mkdirDashP(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		die("Failed to create dir (%s), err (%v)", dir, err) // Haven't seen this happen
	}
}

// Must be run in project's base dir
func initNewProj() {
	// First make sure there's no files
	files, err := ioutil.ReadDir(".")
	if err != nil {
		die("Failed to read current directory (%v)", err)
	}

	for _, f := range files {
		if f.Name() == ".git" && f.IsDir() {
			continue // Only successful case!
		}
		die("Init command found a file or dir (%s). Base dir must be empty save .git", f.Name())
	}

	// Then create some dirs
	mkdirDashP("static/css")
	mkdirDashP("pages/demo")
	mkdirDashP("drafts/demo")

	// Then write em all!
	for fileName, contents := range initialFiles {
		err = ioutil.WriteFile(fileName, []byte(contents), 0644)
		if err != nil {
			die("Couldn't write file (%s), err: (%v)", fileName, err)
		}
	}
}

// ***** Standard project and publishing related files
var gitIgnore = ``
var robots = ``

// TODO support this and other forms like google news, yahoo, etc.
// RSS and Sitemap will need to go into the generation process
var rss = ``
var sitemapXML = ``

// ***** Base HTML template that uses Bootstrap4
var baseHTML = ``

// ***** Example CSS file. You could put js/ img/ in static as well
var mainCSS = ``

// ***** Example pages and drafts
var pagesIndexMD = ``
var pagesDemoIndexMD = ``
var pagesDemoHelpHTML = ``

var draftsIndexHTML = ``
var draftsDemoIndexMD = ``
var draftsDemoAboutHTML = ``

// ***** App Engine hosting related files
var appYaml = ``

// In order to make this work on App Engine, you
// need to have a dynamic server setup even if does
// nothing. TODO link to a blog post about this.
var mainGo = ``
