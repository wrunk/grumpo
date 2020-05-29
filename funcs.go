package main

import (
	"sort"
	"time"
)

// Currently used to hold template rendering functions

// Gives the list of latest pages, newest on top with a limit
// num. Uses latestPages as a cache
func fnLatestPages(num int) []Page {
	if len(latestPages) == 0 {
		latestPages = make([]Page, len(pages))
		copy(latestPages, pages)
		// Sort descending
		sort.Slice(latestPages, func(i, j int) bool {
			return time.Time(latestPages[i].Meta.PublishDate).After(time.Time(latestPages[j].Meta.PublishDate))

		})
	}
	if len(latestPages) <= num {
		return latestPages
	}
	return latestPages[:num]
}
