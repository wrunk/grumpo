package main

import "testing"

type tbp struct {
	Dir string
	Fn  string

	Page
}

// TODO add more test cases and tests in general!
func TestBuildPage(t *testing.T) {
	tableTests := []tbp{{
		Dir: "pages",
		Fn:  "index.html",
		Page: Page{
			BaseDir:       "pages",
			LinkDir:       "",
			FileName:      "index.html",
			Name:          "index",
			Ext:           "html",
			FullPath:      "pages/index.html",
			BuildDir:      "build",
			BuildFullPath: "build/index.html",
		}},
		{
			Dir: "pages/warren",
			Fn:  "author.md",
			Page: Page{
				BaseDir:       "pages/warren",
				LinkDir:       "warren/author",
				FileName:      "author.md",
				Name:          "author",
				Ext:           "md",
				FullPath:      "pages/warren/author.md",
				BuildDir:      "build/warren/author",
				BuildFullPath: "build/warren/author/index.html",
			}},
	}
	for _, tt := range tableTests {
		p := buildPage(tt.Dir, tt.Fn)
		if p.BaseDir != tt.Page.BaseDir {
			t.Errorf("BaseDir got (%s) wanted (%s)", p.BaseDir, tt.Page.BaseDir)
		}
		if p.LinkDir != tt.Page.LinkDir {
			t.Errorf("LinkDir got (%s) wanted (%s)", p.LinkDir, tt.Page.LinkDir)
		}
		if p.FileName != tt.Page.FileName {
			t.Errorf("FileName got (%s) wanted (%s)", p.FileName, tt.Page.FileName)
		}
		if p.Name != tt.Page.Name {
			t.Errorf("Name got (%s) wanted (%s)", p.Name, tt.Page.Name)
		}
		if p.Ext != tt.Page.Ext {
			t.Errorf("Ext got (%s) wanted (%s)", p.Ext, tt.Page.Ext)
		}
		if p.FullPath != tt.Page.FullPath {
			t.Errorf("FullPath got (%s) wanted (%s)", p.FullPath, tt.Page.FullPath)
		}
		if p.BuildDir != tt.Page.BuildDir {
			t.Errorf("BuildDir got (%s) wanted (%s)", p.BuildDir, tt.Page.BuildDir)
		}
		if p.BuildFullPath != tt.Page.BuildFullPath {
			t.Errorf("BuildFullPath got (%s) wanted (%s)", p.BuildFullPath, tt.Page.BuildFullPath)
		}
	}

}
