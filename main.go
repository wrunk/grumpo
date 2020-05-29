package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	baseFile    = "base.html"
	extHTML     = "html"
	extMarkdown = "md"
	version     = 1 // Arbitrary right now
)

var (
	baseTemplate *template.Template
	validExts    = []string{extHTML, extMarkdown}
	pages        = []Page{}
	latestPages  []Page // Cached
)

var help = `Please use grumpo like:
grumpo
	To see the version and this message
	
grumpo <cmd> <space separated flags>
	Where cmd could be: gen|init|local
	Flags could be -nohtmlpretty|-nohtmlvalidate
`

func die(f string, args ...interface{}) {
	fmt.Printf("Error: "+f+"\n", args...)
	os.Exit(1)
}

func helpDie() {
	fmt.Printf(help)
	os.Exit(0)
}

// TODO create init command that will require an empty starting dir
func main() {

	if len(os.Args) == 1 { // Running just grumpo
		fmt.Printf("Grumpo version %d\n", version)
		helpDie()
	}

	cmd := os.Args[1]
	switch cmd {
	case "gen":
		loadBaseTemplate()
		gen()
	case "init":
		initNewProj()
	case "local":
		local()
	default:
		helpDie()
	}
}
