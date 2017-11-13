package main

// “Given a sorted list arr[] find the number of occurrences of a number.
//  ”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

// Brute force: Time Complexity is o(n); Space Complexity is o(1)
func bruteCount(data []int, value int) (count int) {
	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == value {
			count++
		}
	}
	return
}

// binary approach: o(logn); Space Complexity is o(1)
func binaryCount(data []int, value int) int {
	size := len(data)
	var firstIndex, lastIndex func(list []int, a, b, c int) int

	firstIndex = func(d []int, start, end, v int) int {
		if end < start {
			return -1
		}
		mid := (start + end) / 2
		if v == d[mid] && (mid == start || v != d[mid-1]) {
			return mid
		}
		if v < d[mid] {
			return firstIndex(d, start, mid-1, v)
		}
		return firstIndex(d, mid+1, end, v)
	}

	lastIndex = func(d []int, start, end, v int) int {
		if end < start {
			return -1
		}
		mid := (start + end) / 2
		if v == d[mid] && (mid == end || v != d[mid+1]) {
			return mid
		}
		if v < d[mid] {
			return lastIndex(d, start, mid-1, v)
		}
		return lastIndex(d, mid+1, end, v)
	}
	return lastIndex(data, 0, size-1, value) - firstIndex(data, 0, size-1, value)
}
