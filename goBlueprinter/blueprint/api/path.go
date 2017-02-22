package main

import "strings"

// PathSeparator constant
const PathSeparator = "/"

// Path match path and ID
type Path struct {
	Path string
	ID   string
}

// NewPath factory
func NewPath(p string) *Path {
	var id string
	p = strings.Trim(p, PathSeparator)
	s := strings.Split(p, PathSeparator)
	if len(s) > 1 {
		id = s[len(s)-1]
		p = strings.Join(s[:len(s)-1], PathSeparator)
	}
	return &Path{Path: p, ID: id}
}

// HasID doesnt need a comment
func (p *Path) HasID() bool {
	return len(p.ID) > 0
}
