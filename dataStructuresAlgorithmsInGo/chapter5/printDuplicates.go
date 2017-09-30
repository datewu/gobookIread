package main

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/set"
)

// Exhaustive search or Brute force: Time complexity is o(n2); Space complexity is o(1)
func bruteForeRepeating(data []int) {
	size := len(data)
	fmt.Println("Repeating elements are:")

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if data[i] == data[j] {
				fmt.Print(" ", data[i])
			}
		}
	}
}

// sort all elements before search: Time complexity is o(nlogn); Space complexity is o(1)
func afterSortedRepeating(data []int) {
	size := len(data)
	sort.Ints(data)
	fmt.Println("Repeating elements are:")

	for j := 1; j < size; j++ {
		if data[j] == data[j-1] {
			fmt.Print(" ", data[j])
		}
	}
}

// using Hash-Table: Time complexity is o(n); Space complexity is o(n)
func hashTableRepeating(data []int) {
	size := len(data)
	s := set.New()
	fmt.Println("Repeating elements are:")

	for i := 0; i < size; i++ {
		if s.Has(data[i]) {
			fmt.Print(" ", data[i])
		} else {
			s.Insert(data[i])
		}
	}
}

// Counting: Time complexity is o(n); Space complexity is o(n)
func countingRepeating(data []int, intrange int) {
	size := len(data)
	count := make([]int, intrange)
	fmt.Println("Repeating elements are:")

	for i := 0; i < size; i++ {
		if count[data[i]] == 1 {
			fmt.Print(" ", data[i])
		} else {
			count[data[i]]++
		}
	}
}
