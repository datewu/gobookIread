package search

import (
	"fmt"
	"log"
)

// Result contains the result of a search.
type Result struct {
	Filed   string
	Content string
}

// Matcher defines the behavior required by types that want
// to implement a new search type.
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(m Matcher, f *Feed, searchTerm string, r chan<- *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := m.Search(f, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		r <- result
	}
}

// Display writes results to the terminal window as they
// are received by the individual goroutines.
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	for r := range results {
		fmt.Printf("%s: \n%s\n\n", r.Field, r.Content)
	}
}
