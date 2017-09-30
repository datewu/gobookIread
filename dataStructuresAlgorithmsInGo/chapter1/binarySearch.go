package main

import (
	"sort"
)

func binarySearch(data []int, value int) bool {
	sort.Ints(data)

	size := len(data)
	var mid int
	low := 0
	high := size - 1

	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		}
		if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false

}

//
// func main() {
// 	s := []int{1, 5, 76, 8, 9, 4, 9, 8, 3, 5, 67, 4}
// 	fmt.Println(s)
// 	fmt.Println("contains", 9, binarySearch(s, 9))
// 	fmt.Println("contains", 100, binarySearch(s, 100))
// }
