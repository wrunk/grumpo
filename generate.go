package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func gen() {
	// Walk files is currently recursive and cant operate on the
	// global pages. TODO fix this
	pgs, errMsg := walkFiles("pages")
	if errMsg != "" {
		die(errMsg)
	}
	pages = pgs

	errMsg = validateIndex()
	if errMsg != "" {
		die(errMsg)
	}
	generateAndWriteHTML()
	errMsg = copyStatic()
	if errMsg != "" {
		die(errMsg)
	}
}

func generateAndWriteHTML() {
	// Loop over all pages in the index:
	for _, p := range pages {
		// Make this file's base dir (no error if already exists,
		// same behavior as mkdir -p)
		err := os.MkdirAll(p.BuildDir, 0755)
		if err != nil {
			die("Failed to create dir %s", err) // Haven't seen this happen
		}
		pageDataBys, err := ioutil.ReadFile(p.FullPath)
		if err != nil {
			die("Couldn't read file %s %s", p.FullPath, err)
		}
		if p.Ext == extMarkdown {
			pageDataBys = buildMarkdown(pageDataBys)
		}
		// Render full page
		finalPage := renderHTML(map[string]interface{}{
			"page":  string(pageDataBys),
			"pages": pages, // Pass in all pages for future use
		})

		err = validateHTML(finalPage)
		if err != nil {
			die("%s Resulted in invalid html (%s)", p.FullPath, err)
		}

		// Pretty print after we validate since this stupid lib won't check crap!
		// Don't use this for now as it screws up the pre formatted code blocks
		// finalPage = gohtml.FormatBytes(finalPage)
		err = ioutil.WriteFile(p.BuildFullPath, finalPage, 0644)
		if err != nil {
			die("Couldn't write file (%s), err: (%v)", p.BuildFullPath, err)
		}
	}
}

func copyStatic() string {
	// This is lazy but leverage system cp. Will also create dir
	err := exec.Command("cp", "-r", "./static", "build").Run()
	if err != nil {
		return fmt.Sprintf("Failed to copyStatic using system cp cmd: (%v)", err)
	}
	return ""
}
