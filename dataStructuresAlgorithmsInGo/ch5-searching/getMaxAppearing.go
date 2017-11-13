package main

import "sort"

// Exhaustive search or Brute force: Time Complexity is o(n2); Space Complexity is o(1)
func bruteMax(data []int) int {
	size := len(data)
	max := data[0]
	maxCount := 1

	for i := 0; i < size; i++ {
		count := 1
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
	return max
}

// sort all the element: Time Complexity is o(nlogn); Space Complexity is o(1)
func sortMax(data []int) int {
	size := len(data)
	max := data[0]
	maxCount := 1
	curr := data[0]
	currCount := 1
	sort.Ints(data)

	for i := 1; i < size; i++ {
		if data[i] == data[i-1] {
			currCount++
		} else {
			currCount = 1
			curr = data[i]
		}
		if currCount > maxCount {
			maxCount = currCount
			max = curr
		}
	}
	return max
}

// counting, this approach is only possible if we know the range of the input: Time Complexity is o(n), Space Complexity is o(n)
func countMax(data []int, dataRange int) int {
	max := data[0]
	maxCount := 1
	size := len(data)
	count := make([]int, dataRange)

	for i := 0; i < size; i++ {
		count[data[i]]++
		if count[data[i]] > maxCount {
			maxCount = count[data[i]]
			max = data[i]
		}
	}
	return max
}
