package main

// “Selection-Sort searches the whole unsorted array and put the largest value at the end of it. This algorithm is having the same Time Complexity, but performs better than both bubble and Insertion-Sort as less number of comparisons required. The sorted list is created backward in Selection-Sort.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

func selectSort(data []int) {
	size := len(data)
	var i, j, maxIndex int
	for i = 0; i < size-1; i++ {
		maxIndex = 0
		for j = 1; j < size-i-1; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[size-i-1], data[maxIndex] = data[maxIndex], data[size-i-1]
	}

}
