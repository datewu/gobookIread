package thesaurus

// Thesaurus interface just descibe Synonyms method
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
