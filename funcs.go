package main

import "sort"

// Currently used to hold template rendering functions

// Gives the list of latest pages, newest on top with a limit
// num. Uses latestPages as a cache
func fnLatestPages(num int) []Page {
	if len(latestPages) == 0 {
		latestPages = make([]Page, len(pages))
		copy(latestPages, pages)
		// Sort descending
		sort.Slice(latestPages, func(i, j int) bool { return latestPages[i].Meta.PublishDate.After(latestPages[j].Meta.PublishDate) })
	}
	if len(latestPages) <= num {
		return latestPages
	}
	return latestPages[:num]
}
