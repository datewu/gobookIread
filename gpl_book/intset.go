package main

import (
	"bytes"
	"fmt"
)

func main() {
	setNew := new(IntSet)
	data := []int{2, 2, 3, 3, 312, 35, 67, 97, 10, 25}
	seedSet(setNew, data)
	fmt.Println(setNew, setNew.Has(13))
	fmt.Println(setNew.Has(25), setNew.Has(13))
}

func seedSet(s *IntSet, d []int) {
	for _, v := range d {
		s.Add(v)
	}
}

// IntSet is a set of small non-negative integers.
// Its zero calue represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	// if s.Has(x) {
	// 	return
	// }
	word, bit := x/64, uint(x%64)
	if word < len(s.words) && s.words[word]&(1<<bit) != 0 {
		return
	}
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String return the set  as a string of the form "{ 1 5 9 }".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
