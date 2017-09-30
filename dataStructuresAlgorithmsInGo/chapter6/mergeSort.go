package main

// “Merge sort divide the input into half recursive in each step. It sort the two parts separately recursively and finally combine the result into final sorted output.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func mergeSort(data []int, comp func(int, int) bool) {
	size := len(data)
	tempArray := make([]int, size)

	//	var merge func([]int, []int, int, int, int, func(int, int) bool)
	merge := func(a []int, tempA []int, lowerIndex, middleIndex, upperIndex int, c func(int, int) bool) {
		lowerStart, lowerStop, upperStart, upperStop := lowerIndex, middleIndex, middleIndex+1, upperIndex
		count := lowerIndex
		for lowerStart <= lowerStop && upperStart <= upperStop {
			if !c(a[lowerStart], a[upperStart]) {
				tempA[count] = a[lowerStart]
				lowerStart++
			} else {
				tempA[count] = a[upperStart]
				upperStart++
			}
			count++
		}
		for lowerStart <= lowerStop {
			tempA[count] = a[lowerStart]
			count++
			lowerStart++
		}
		for upperStart <= upperStop {
			tempA[count] = a[upperStart]
			count++
			upperStart++
		}
		for i := lowerIndex; i <= upperIndex; i++ {
			a[i] = tempA[i]
		}
	}

	var mergeRecursie func([]int, []int, int, int, func(int, int) bool)

	mergeRecursie = func(arr []int, temp []int, lIndex, hIndex int, com func(int, int) bool) {
		if lIndex >= hIndex {
			return
		}
		middleIndex := (lIndex + hIndex) / 2
		mergeRecursie(arr, temp, lIndex, middleIndex, com)
		mergeRecursie(arr, temp, middleIndex+1, hIndex, com)
		merge(arr, temp, lIndex, middleIndex, hIndex, com)
	}
	mergeRecursie(data, tempArray, 0, size-1, comp)
}

// func main() {
// 	data := []int{9, 3, 8, 1, 56, 2, 9, 54, 12, 4, 67, 24, 79}
// 	fmt.Println(data)
// 	mergeSort(data, func(a int, b int) bool {
// 		return a > b
// 	})
// 	fmt.Println(data)
// }
