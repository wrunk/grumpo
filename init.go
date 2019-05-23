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
		"robots.txt":              robots,
		"rss.txt":                 rss,
		"sitemap.xml":         sitemapXML,
		"base.html":           baseHTML,
		"static/css/main.css": mainCSS,

		"pages/index.md":       pagesIndexMD,
		"pages/demo/index.md":  pagesDemoIndexMD,
		"pages/demo/help.html": pagesDemoHelpHTML,

		"drafts/index.html":      draftsIndexHTML,
		"drafts/demo/index.md":   draftsDemoIndexMD,
		"drafts/demo/about.html": draftsDemoAboutHTML,

		"app.yaml": appYaml,
		"main.go":  mainGo,
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
var gitIgnore = `
# Shouldnt need too many ignores for a static site
.idea
.vscode
.DS_Store
`

// Robots txt tells crawlers like GoogleBot what not to crawl
// Generally this is for gated sections like /users/ or similar
// See https://www.nytimes.com/robots.txt for example.
// If you wanted to disable all robots you could:
// User-agent: *
// Disallow: /
//
// Of course no robots or humans are forced to obey this :)
var robots = `
Sitemap: /sitemap.xml
`

// TODO support this and other forms like google news, yahoo, etc.
// RSS and Sitemap will need to go into the generation process
var rss = ``
var sitemapXML = ``

// ***** Base HTML template that uses Bootstrap4
// This is "Starter Template" from bootstrap 4 at
//
var baseHTML = `
<!doctype html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" 
		integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
		
		<!-- Add our own CSS -->
		<link rel="stylesheet" href="/static/css/main.css">

    <title>Grumpo, Example</title>
  </head>
  <body>
		<h1>Grumpo, Example</h1>
		<p>You can check out Grumpo <a href="https://github.com/wrunk/grumpo">here.</a></p>

		<p>Page content goes below which is usually wrapped in bootstrap containers/row/cols</p>

		{{ .page }}

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
  </body>
</html>
`

// ***** Example CSS file. You could put js/ img/ in static as well
var mainCSS = `
/* Example CSS */
h1 {
	text-decoration: underline;
}
`

// ***** Example pages and drafts
var pagesIndexMD = `
# Home Page!

This part should be a paragraph about the home page!
Check the html to confirm.
`
var pagesDemoIndexMD = `
# Demo Index Markdown Page

Just showing use of a subdir...
`
var pagesDemoHelpHTML = `
<h1>Pages can also be written in html</h1>

<p>This demo help page is html!</p>
`

var draftsIndexHTML = `
<h1>Home Page in progress</h1>

<p>We leverage a simple drafts directory to deal with
drafts that you don't want to get built and go live</p>`

var draftsDemoIndexMD = `
# Demo Index Markdown page

Back into markdown. Hopefully we can do more cool stuff
with markdown like TOC and auto heading anchoring.
`

var draftsDemoAboutHTML = `
<h1>Drafts Demo About HTML Page</h1>

<p>Showing a sub dir in drafts. To complete a draft, move it into pages
wherever you would like it.</p>`

// ***** App Engine hosting related files
// Main app engine hosting configuration file.
// You can see more info about using App Engine with Golang here:
// https://cloud.google.com/appengine/docs/standard/go112/quickstart
// However you need careful settings to serve a static site
// TODO add comments or a post on this
var appYaml = `
runtime: go111
service: grumpo
env: standard
instance_class: F1

# Make sure our content isn't cached too long (10m is default)
default_expiration: "60s"

handlers:
- url: /
  static_files: build/index.html
  upload: /
  secure: always

# TODO add rules for sitemaps, robots, rss

- url: /static
  static_dir: build/static
  secure: always

- url: /(.*)/$
  static_files: build/\1/index.html
  upload: .*\.html$
  secure: always

- url: '.*'
  script: auto

# Do we need this?
automatic_scaling:
  min_idle_instances: automatic
  max_idle_instances: automatic
  min_pending_latency: automatic
  max_pending_latency: automatic

`

// In order to make this work on App Engine, you
// need to have a dynamic server setup even if does
// nothing. TODO link to a blog post about this.
// For now any requests that end up here are assumed to be 404.
// TODO consider how to link a 404.html in here
var mainGo = `
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// Always return 404 for now since this is a static site
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

`
