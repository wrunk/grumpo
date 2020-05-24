package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// TODO this file could be refactored out into proper handler
// functions that properly set the content type AND use a basic
// middleware or setup function to handle the common stuff

func local() {
	fmt.Println("Running a local web server on port 9876")
	fmt.Println("Browse to http://localhost:9876/__index__ to view the index")

	// All requests other than index and json
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		loadBaseTemplate()
		errMsg := loadAllFromDisk()
		if errMsg != "" {
			writeErrorPage(w, "Error loading pages", fmt.Sprintf("Error: (%s)", errMsg))
			return
		}
		// This is a hacked version of loading static files as it
		// would exist on the server
		path := r.URL.Path[1:]
		page := findPage(path)
		if page == nil {
			writeErrorPage(w, "Page Not Found", fmt.Sprintf("Page (%s) was not found", path))
			return
		}
		fmt.Fprintf(w, string(renderPage(*page)))
	})

	http.HandleFunc("/__index__", func(w http.ResponseWriter, r *http.Request) {
		errMsg := loadAllFromDisk()
		if errMsg != "" {
			writeErrorPage(w, "Error loading pages", fmt.Sprintf("Error: (%s)", errMsg))
			return
		}

		fmt.Fprintf(w, renderLocalIndex())
	})

	http.HandleFunc("/__json__", func(w http.ResponseWriter, r *http.Request) {
		errMsg := loadAllFromDisk()
		if errMsg != "" {
			writeErrorPage(w, "Error loading pages", fmt.Sprintf("Error: (%s)", errMsg))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, renderJSONRepr())
	})

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", fs)

	http.ListenAndServe(":9876", nil)
}

func writeErrorPage(w http.ResponseWriter, errTitle, errMsg string) {
	fmt.Fprintf(w, "<body><h1>%s</h1><p>%s</p></body>", errTitle, errMsg)
}

// Although not exactly how a webserver will look for the static files
// on disk, easy enough to just find on linkdir
func findPage(path string) *Page {
	path = strings.TrimSuffix(path, "/")
	for i, page := range pages {
		if path == page.LinkDir {
			return &(pages[i]) // Don't return &page as that memory is reused
		}
	}
	return nil
}

func renderLocalIndex() string {
	page := "<body>\n<h1>Local Development Index</h1>\n<h2>Published Pages</h2>\n"
	for _, p := range pages {
		page += localLink(p)
	}
	jsonSection := `
	<h2>JSON Repr of Internal Pages</h2>
	<p><a href="/__json__">JSON Repr</a></p>
	`
	return page + "\n" + jsonSection + "</body>"
}

func renderJSONRepr() string {
	bys, err := json.Marshal(pages)
	// Almost no chance of error ever so just panic
	if err != nil {
		panic(err)
	}
	return string(bys)
}

func localLinkHTML(url string) string {
	return fmt.Sprintf(`<div><a href="%s">%s</a></div>`, url, url) + "\n"
}

func localLink(page Page) string {
	u := "http://localhost:9876/"
	if page.LinkDir == "" { // will only happen for home page
		return localLinkHTML(u)
	}
	return localLinkHTML(u + page.LinkDir + "/")
}

// Find all pages from disk to populate pages var
func loadAllFromDisk() string {
	pgs, errMsg := walkFiles("pages")
	if errMsg != "" {
		return errMsg
	}
	pages = pgs
	return validateIndex()
}

// TODO change this to using error messages not dying so we can
// use with web server
// TODO combine with generate() below
func renderPage(page Page) []byte {
	pageDataBys, err := ioutil.ReadFile(page.FullPath)
	if err != nil {
		die("Couldn't read file %s %s", page.FullPath, err)
	}
	if page.Ext == extMarkdown {
		pageDataBys = buildMarkdown(pageDataBys)
	}
	// Render full page!
	finalPage := renderHTML(map[string]interface{}{"page": string(pageDataBys)})

	err = validateHTML(finalPage)
	if err != nil {
		die("%s Resulted in invalid html (%s)", page.FullPath, err)
	}

	// Pretty print after we validate since this stupid lib won't check crap!
	// Don't use this for now as it screws up the pre formatted code blocks
	// finalPage = gohtml.FormatBytes(finalPage)
	return finalPage
}
