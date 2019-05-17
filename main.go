package main

import (
	"fmt"
	"os"
	"text/template"
)

var baseTemplate *template.Template

const (
	baseFile    = "base.html"
	extHTML     = "html"
	extMarkdown = "md"
	version     = 1 // Arbitrary right now
)

var (
	validExts = []string{extHTML, extMarkdown}
	pages     = []Page{}
)

func die(f string, args ...interface{}) {
	fmt.Printf("Error: "+f+"\n", args...)
	os.Exit(1)
}

var help = `Please use grumpo like:
grumpo
	To see the version and this message
	
grumpo <cmd> <space separated flags>
	Where cmd could be: gen|init|local
	Flags could be -nohtmlpretty|-nohtmlvalidate
`

func helpDie() {
	fmt.Printf(help)
	os.Exit(0)
}

// TODO create init command that will require an empty starting dir
func main() {
	loadBaseTemplate()

	if len(os.Args) == 1 { // Running just grumpo
		fmt.Printf("Grumpo version %d\n", version)
		helpDie()
	}

	cmd := os.Args[1]
	switch cmd {
	case "gen":
		gen()
	case "init":
		initNewProj()
	case "local":
		local()
	default:
		helpDie()
	}

}
