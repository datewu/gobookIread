package main

import "fmt"

// “Given a list of n-1 elements, which are in the range of 1 to n. There are no duplicates in the list. One of the integer is missing. Find the missing element.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

// Brute force: Time Complexity is o(n2); Space Complexity is o(1)
func bruteMissing(data []int) (int, bool) {
	var missing bool
	size := len(data)

	for i := 0; i < size; i++ {
		missing = true
		for j := 0; j < size; j++ {
			if data[j] == i {
				missing = false
				break
			}
		}
		if missing {
			return i, true
		}

	}
	fmt.Println("There is no number missing")
	return 0, false
}

// xor : Time Complexity is o(n); Space Complexity is o(1)
func xorMissine(data []int, dataRange int) (int, bool) {
	size := len(data)
	xorSum := 0

	for i := 1; i <= dataRange; i++ {
		xorSum ^= i
	}
	for i := 0; i < size; i++ {
		xorSum ^= data[i]
	}
	if xorSum == 0 {
		fmt.Println("There is no number missing")
		return 0, false
	}
	return xorSum, true
}
