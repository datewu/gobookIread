package main

// “Insertion-Sort Time Complexity is O(n2) which is same as Bubble-Sort but perform a bit better than it. It is the way we arrange our playing cards. We keep a sorted subarray. Each value is inserted into its proper position in the sorted sub-array in the left of it.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

func insertSort(data []int, comp func(int, int) bool) []int {
	size := len(data)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = data[i]
		for j = i; j > 0 && comp(data[j-1], temp); j-- {
			data[j] = data[j-1]
		}
		data[j] = temp
	}
	return data
}
