package main

import (
	"fmt"
	"sort"
)

// “Given a list of n elements. Find the majority element, which appears more than n/2 times. Return 0 in case there is no majority element.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

// Exhaustiv search or brute force: Time Complexity is o(n2); Space Complexity is o(1)
func bruteMajority(data []int) (int, bool) {
	size := len(data)
	max := 0
	count := 0
	maxCount := 0

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if data[i] == data[j] {
				count++
			}
		}
		if count > maxCount {
			max = data[i]
			maxCount = count
		}
	}
	if maxCount > size/2 {
		return max, true
	}
	fmt.Println("Majority element does not exist")
	return 0, false
}

// sort: Time Complexity is o(nlogn); Space Complexity is o(1)
func sortMajority(data []int) (int, bool) {
	size := len(data)
	majIndex := size / 2
	sort.Ints(data)
	candidate := data[majIndex]
	count := 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return candidate, true
	}
	fmt.Println("Majority element does not exist")
	return 0, false
}

// cancelation approach: Time Complexity is o(n); Space Complexity is o(1)
func cancelMajority(data []int) (int, bool) {
	size := len(data)
	majIndex := 0
	count := 0

	for i := 0; i < size; i++ {
		if data[i] == data[majIndex] {
			count++
		} else {
			count--
		}
		if count == 0 {
			majIndex = i
			count = 1
		}
	}
	candidate := data[majIndex]
	count = 0
	for i := 0; i < size; i++ {
		if data[i] == candidate {
			count++
		}
	}
	if count > size/2 {
		return candidate, true
	}
	fmt.Println("Majority element does not exist")
	return 0, false
}
