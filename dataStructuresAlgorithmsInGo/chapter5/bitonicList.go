package main

import "fmt"

// “A bitonic list comprises of an increasing sequence of integers immediately followed by a decreasing sequence of integers.”
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func searchMax(data []int) (int, bool) {
	size := len(data)
	if size < 3 {
		fmt.Println("Invalid Input")
		return 0, false
	}

	start, end := 0, size-1
	maxFound := 0

	for start <= end {
		mid := (start + end) / 2
		if data[mid-1] < data[mid] && data[mid+1] < data[mid] {
			maxFound = 1
			break
		} else if data[mid-1] < data[mid] && data[mid] < data[mid+1] {
			// increasing
			start = mid + 1
		} else if data[mid-1] > data[mid] && data[mid] > data[mid+1] {
			// decreasing
			end = mid - 1
		} else {
			break
		}
	}
	if maxFound == 0 {
		fmt.Println("No maxima Found")
		return 0, false
	}

	return maxFound, true
}

func serachValue(data []int, v int) int {
	size := len(data)
	maxIndex, _ := searchMax(data)

	k := customBinarySearch(data, 0, maxIndex, v, true)
	if k != -1 {
		return k
	}
	return customBinarySearch(data, maxIndex+1, size-1, v, false)
}

func customBinarySearch(data []int, start, end, value int, isInc bool) int {
	if end < start {
		return -1
	}
	mid := (start + end) / 2
	if value == data[mid] {
		return mid
	}
	if isInc && value < data[mid] || isInc == false && value > data[mid] {
		return customBinarySearch(data, start, mid-1, value, isInc)
	}
	return customBinarySearch(data, mid+1, end, value, isInc)
}
