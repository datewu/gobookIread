package main

import "fmt"
import "sort"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"database":              {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, course)
	}
}

func topSort(m map[string][]string) (ordered []string) {
	seen := make(map[string]bool)

	var travelAll func([]string)
	travelAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				travelAll(m[item])
				//fmt.Println(item)
				ordered = append(ordered, item)
			}

		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	travelAll(keys)
	return
}
