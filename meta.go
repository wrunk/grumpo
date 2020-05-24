package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Metadata taken from the top of the article if available
// See readme for metadata details
type Meta struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
	Headline1   string `json:"hl1"`
	Headline2   string `json:"hl2"` // Can be used for tweets or a/b
	Headline3   string `json:"hl3"`
	Image       string `json:"image"`
	ImageAlt    string `json:"image_alt"`

	// The publish date to show on the article and elsewhere.
	// Controls if this shows up in things like recent posts
	// TODO throw an error if an article is published but not live
	PublishDate localDate `json:"publish_date"`

	// If live is set to false (default) the article will NOT be generated
	// for deployment or show up in sitemaps or rss
	// On local we should really show a banner notifying that this article is not
	// live AND we should check all links to make sure it isnt linked anywhere
	// Article can be set live before before it is published so you can
	// preview it on the live website but it wont show up in recent posts
	// until it is published
	Live bool `json:"live"`

	// If set, will show an updated at field below original publication date
	UpdatedDate localDate `json:"updated_date"` // Could be set from git as well

	AuthorID int `json:"author_id"` // Simple author ID linked to authors.yaml

	// AKA don't use the base template
	SkipBaseTemplate bool `json:"skip_base_template"`

	// If the json metadata ends on line 10, then content start on 11
	contentStartsOn int
}

func (m *Meta) Validate() error {
	if m.Title == "" {
		return fmt.Errorf("Meta must have a title")
	}
	n := time.Now()
	pd := time.Time(m.PublishDate)
	// We've set the pub date to the past but the article isn't considered live
	// which doesn't make sense.
	if !pd.IsZero() && n.After(pd) && !m.Live {
		return fmt.Errorf("Publish date cannot be in the past without live being true")
	}
	return nil
}

// Basic, custom datetime that supports either:
// 2020-01-02 OR
// 2020-01-02:00:00:00
// AND NOTHING ELSE! ALL TIMES LOCAL! If you build grumpo
// on a server, make sure the OS is set to your publishing timezone
type localDate time.Time

// imeplement Marshaler und Unmarshalere interface
func (ld *localDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" {
		return nil
	}
	// Try format 1
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		// Didn't work, try format 2
		t, err = time.Parse("2006-01-02:03:04:05", s)
		if err != nil {
			return err
		}
	}
	*ld = localDate(t)
	return nil
}

/*
loadMeta:
- Tries to open provided filePath
- Reads line by line, doing a very rough json parse
- If we find something valid actually try to unmarshal into Meta
- Validate Meta
- Return any errors
- Return the Meta object
*/
func loadMeta(filePath string) (*Meta, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var opens, closes, lineNum, contentStarts int
	jsonBody := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") { //Ignore empty or comment lines
			continue
		}
		opens += strings.Count(line, "{")
		closes += strings.Count(line, "}")
		if opens > 0 {
			jsonBody += line
		}
		if opens > 0 && opens == closes {
			contentStarts = lineNum + 1 // Save where the content starts
			break
		}
	}
	if jsonBody == "" {
		return nil, fmt.Errorf("Found no json ({...})(%s)", filePath)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	m := &Meta{}
	err = json.Unmarshal([]byte(jsonBody), m)
	if err != nil {
		return nil, err
	}

	err = m.Validate()
	if err != nil {
		return nil, fmt.Errorf("Failed to load meta for article (%s) (%s)", filePath, err.Error())
	}
	m.contentStartsOn = contentStarts
	return m, nil
}
