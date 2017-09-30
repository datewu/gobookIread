package main

// “Find median of two sorted Lists.
// First approach: Keep track of the index of both the list, say the index are i and j. keep increasing the index of the list which ever have a smaller value. Use a counter to keep track of the elements that we have already traced.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

func findMedian(dataOne []int, dataTwo []int) int {
	sizeOne, sizeTwo := len(dataOne), len(dataTwo)
	medianIndex := ((sizeOne+sizeTwo)%2 + sizeOne + sizeTwo) / 2

	var i, j, count int
	for ; count < medianIndex-1; count++ {
		if i < sizeOne-1 && dataOne[i] < dataTwo[j] {
			i++

		} else {
			j++
		}
	}
	if dataOne[i] < dataTwo[j] {
		return dataOne[i]
	}
	return dataTwo[j]
}
