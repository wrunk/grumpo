package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLoadMeta(t *testing.T) {

	// TODO improve all this to use table tests
	fn := ".testfile.md"
	ioutil.WriteFile(fn, []byte(`
	{
		// You can use c-style comments as long as they are on their own
		// line and use the // format (not /* */)
	  
		// Title is ALWAYS required. You can control how this is used in
		// base.html
		"title": "Great Blog Post Aboot Cats!",
	  
		// Used in OG tags and can be used in recent posts promo
		"desc": "Here is a short(ish) description about this article",
	  
		// Headlines can be used in various ways. More on this later
		"hl1": "Headlines could be variations on title for",
		"hl2": "a/b testing purposes or",
		"hl3": "auto tweeting so you post your content multiple times spread out",
	  
		// Specify a relative or full URL to the article's canonical image.
		// We'll set article's og:image tag to this
		"image": "/static/img/hello.jpg",
		"image_alt": "A hello face",
	  
		// If set to true, this page will go live with grumpo gen
		// However it wont show up in recent posts until you set
		// a publish date
		"live": false,
	  
		// Both support the format 2020-01-01 OR 2020-01-01:03:04:05
		// No other formats are supported and no timezones can be passed in
		// grumpo commands use your machine's local time zone
		"publish_date": "2050-01-01",
		"updated_date": "2100-01-01",
	  
		// Defaults to false. Set to true to treat this page as a full html page
		"skip_base_template": true
	  }
	  # Some Random Page W/ Markdowns!

Hellos thers somethingslkjsdl falksjf aslkfjas

More shit that would normally be in a real content/MD file
	`), 0777)

	m, err := loadMeta(fn)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Meta: (%s)", spew.Sdump(m))
		if m.contentStartsOn != 37 {
			t.Errorf("Expected content starts on 37, got %d", m.contentStartsOn)
		}
		fileBytes, err := readOffsetFile(fn, 37)
		if err != nil {
			t.Error(err)
		} else if !strings.HasPrefix(strings.TrimSpace(string(fileBytes)), "# Some Random Page") {
			t.Error("didnt find the rest of the file on offset!!")

		}
	}

	err = os.Remove(fn)
	if err != nil {
		t.Error(err)
	}
}
