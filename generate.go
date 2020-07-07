package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func gen() {
	loadBaseTemplate()
	errMsg := loadAllFromDisk()
	if errMsg != "" {
		die(errMsg)
	}

	// TODO fix me
	generateAndWriteHTML()

	// And the rest should work
	errMsg = copyStatic()
	if errMsg != "" {
		die(errMsg)
	}
}

// TODO MUST be combined with local functions!
func generateAndWriteHTML() {
	// Loop over all pages in the index:
	for _, p := range pages {
		if !p.Meta.Live {
			log.Printf("Skipping draft page %s", p.FullPath)
			continue
		}
		// Make this file's base dir (no error if already exists,
		// same behavior as mkdir -p)
		err := os.MkdirAll(p.BuildDir, 0755)
		if err != nil {
			die("Failed to create dir %s", err) // Haven't seen this happen
		}

		finalPage := renderHTML(templateData(p))

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
