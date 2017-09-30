package main

// “Quick sort is also a recursive algorithm.
// ·         In each step, we select a pivot (let us say first element of list).
// ·         Then we traverse the rest of the list and copy all the elements of the list which are smaller than the pivot to the left side of list
// ·         We copy all the elements of the list, which are greater than pivot to the right side of the list. Obviously, the pivot is at its sorted position.
// ·         Then we sort the left and right subarray separately.
// ·         When the algorithm returns the whole list is sorted.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

func quickSort(data []int, comp func(int, int) bool) {
	size := len(data)

	var quickRecursive func([]int, int, int, func(int, int) bool)

	quickRecursive = func(arr []int, lower, upper int, co func(int, int) bool) {
		if lower >= upper {
			return
		}
		pivot := arr[lower]
		start, stop := lower, upper

		for lower < upper {
			for !co(arr[lower], pivot) && lower < upper {
				lower++
			}

			for co(arr[upper], pivot) && lower <= upper {
				upper--
			}

			if lower < upper {
				arr[upper], arr[lower] = arr[lower], arr[upper]
			}

		}
		arr[upper], arr[start] = arr[start], arr[upper] // upper is the pivot position
		quickRecursive(arr, start, upper-1, co)
		quickRecursive(arr, upper+1, stop, co)

	}
	quickRecursive(data, 0, size-1, comp)
}
