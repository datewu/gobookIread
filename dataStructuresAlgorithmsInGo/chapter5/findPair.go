package main

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/set"
)

//
// “Given a list of n numbers, and a value: find two elements such that their sum is equal to “value”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

// Brute force: Time Complexity is o(n2); Space Complexity is o(1)
func brutePair(data []int, value int) bool {
	size := len(data)
	ret := false

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if (data[i] + data[j]) == value {
				fmt.Println("The pair is :", data[i], data[j])
				if !ret {
					ret = true
				}
			}
		}
	}
	return ret
}

// sort: Time Complexity is o(nlogn); Space Complexity is o(1)
func sortPair(data []int, value int) bool {
	size := len(data)
	head := 0
	tail := size - 1
	ret := false
	sort.Ints(data)

	for head < tail {

		curr := data[head] + data[tail]
		if curr == value {
			fmt.Println("The pair is:", data[head]+data[tail])
			ret = true
		}
		if curr < value {
			head++
		} else {
			tail--
		}
	}
	return ret
}

// hash-table: Time Complexity is o(n); Space Complexity is o(n)
func hashPair(data []int, value int) bool {

	s := set.New()
	size := len(data)
	ret := false

	for i := 0; i < size; i++ {
		if s.Has(value - data[i]) {
			fmt.Println("The pair is:", data[i], value-data[i])
			ret = true
		} else {
			s.Insert(data[i])
		}
	}
	return ret
}

//
// “Given two list X and Y. Find a pair of elements (xi, yi) such that xi∈X and yi∈Y where xi+yi=value.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
