package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Recursively walk down dirs finding files
// Since pages var is now global, probably dont need this, whatev
func walkFiles(dir string) ([]Page, string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Sprintf("Failed to read pages dir (%s)", err)
	}
	pgs := []Page{}
	for _, f := range files {
		if f.IsDir() {
			pgs2, errMsg := walkFiles(dir + "/" + f.Name())
			if errMsg != "" {
				return nil, errMsg
			}
			pgs = append(pgs, pgs2...)
		} else {
			pgs = append(pgs, buildPage(dir, f.Name()))
		}
	}
	return pgs, ""
}

/*
buildPage builds the page object with page metadata
Dir possibilities:
- pages (base)
- pages/warren (subfolder)
- pages/warren/trees (sub-subfolder)

File possibilities:
- index.html
- some-file.md
- file_with_no_ext

*/
func buildPage(dir, fileName string) Page {
	p := _buildPage(dir, fileName)
	meta, err := loadMeta(p.FullPath)
	if err != nil {
		die("Failed to load meta from file (%s), err: (%s)", p.FullPath, err)
	}
	p.Meta = meta
	return p
}

// Just construct the Page object so we can test this!
func _buildPage(dir, fileName string) Page {
	// linkDir is used to build links for index page and elsewhere
	linkDir := ""
	// Always exclude /pages/ from links
	if strings.HasPrefix(dir, "pages/") {
		linkDir = dir[6:]
	}
	bdir := strings.Replace(dir, "pages", "build", -1)
	name, ext := getNameAndExt(fileName)
	if name != "index" {
		bdir += "/" + name
		if linkDir != "" { // We're in a sub dir not like about.md
			linkDir += "/" + name
		} else { // In the case of root level pages, need name
			linkDir = name
		}
	}
	return Page{
		BaseDir:       dir,
		LinkDir:       linkDir,
		FileName:      fileName,
		Name:          name,
		Ext:           ext,
		FullPath:      dir + "/" + fileName,
		BuildDir:      bdir,
		BuildFullPath: bdir + "/index.html",
	}
}

func getNameAndExt(fileName string) (string, string) {
	parts := strings.Split(fileName, ".")
	if !strings.Contains(fileName, ".") {
		return fileName, "" // This is really an error case, just return
		// file name
	}
	return strings.Join(parts[:len(parts)-1], "."), parts[len(parts)-1]
}
