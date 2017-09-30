package main

// “Bubble-Sort is the slowest algorithm for sorting. It is easy to implement and used when data is small.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func bubbleSort(data []int, comp func(int, int) bool) {
	size := len(data)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if comp(data[j], data[j+1]) {
				data[j+1], data[j] = data[j], data[j+1]
			}
		}
	}
}

//
// “When there is no more swap in one pass of the outer loop. It indicates that all the elements are already in order so we should stop sorting. This sorting improvement in Bubble-Sort is extremely useful when we know that, except few elements rest of the list is already sorted.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func improvedBubbleSort(data []int, comp func(int, int) bool) {
	size := len(data)
	swapped := 1

	for i := 0; i < size-1 && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < size-i-1; j++ {
			if comp(data[j], data[j+1]) {
				data[j+1], data[j] = data[j], data[j+1]
				swapped = 1
			}
		}
	}
}
