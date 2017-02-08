package main

import (
	"fmt"
	"sort"
)

type stringSlice []string

func (p stringSlice) Len() int {
	return len(p)
}

func (p stringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p stringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var names = []string{
		"da",
		"wang",
		"jiao",
		"wo",
		"lai",
		"xun shan",
		"o",
		"xun wan dong shan",
		"xun xi san ",
		"o",
		"lol dota",
	}
	fmt.Println(names)
	// sort.Sort(stringSlice(names))
	sort.Strings(names)
	fmt.Println(names)

}
