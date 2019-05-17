package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getNameAndExt(fileName string) (string, string) {
	parts := strings.Split(fileName, ".")
	if !strings.Contains(fileName, ".") {
		return fileName, "" // This is really an error case, just return
		// file name
	}
	return strings.Join(parts[:len(parts)-1], "."), parts[len(parts)-1]
}

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
Dir possibilities:
- pages (base)
- pages/warren (subfolder)
- pages/warren/trees (sub-subfolder)

- drafts (in the case of local server)
- drafts/pages (error case - can't have a folder pages)

File possibilities:
- index.html
- some-file.md
- file_with_no_ext

*/
func buildPage(dir, fileName string) Page {
	linkDir := ""
	if strings.HasPrefix(dir, "pages/") {
		linkDir = dir[6:]
	} else if strings.HasPrefix(dir, "drafts") {
		linkDir = dir
	}
	bdir := strings.Replace(dir, "pages", "build", -1)
	name, ext := getNameAndExt(fileName)
	if name != "index" {
		bdir += "/" + name
		linkDir += "/" + name
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
