package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var m defaultMatcher
	Register("default", m)
}

// Search implements the behavior for the default matcher.
func (defaultMatcher) Search(f *Feedm, searchTerm string) ([]*Result, error) {
	return nil, nil
}
