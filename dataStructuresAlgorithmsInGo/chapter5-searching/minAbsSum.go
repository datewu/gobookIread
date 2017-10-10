package main

import (
	"fmt"
	"math"
	"sort"
)

// “Given a List of integers, both +ve and -ve. You need to find the two elements such that their sum is closest to zero.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

// Brute force: Time Complexity is o(n2); Space Complexity is o(1)
func bruteMinAbs(data []int) {
	size := len(data)
	if size < 2 {
		fmt.Println("Invalid Input")
		return
	}

	var sum float64

	minFirst := 0
	minSecond := 1
	minSum := math.Abs(float64(data[0] + data[1]))

	for i := 0; i < size-1; i++ {
		for j := i; j < size; j++ {
			sum = math.Abs(float64(data[i] + data[j]))
			if sum < minSum {
				minSum = sum
				minFirst = i
				minSecond = j
			}
		}
	}
	fmt.Println("The two elements with minimum sum are:", data[minFirst], data[minSecond])
}

// sort: Time Complexity is o(n2); Space Complexity is o(1)
func sortMinAbs(data []int) {
	size := len(data)
	if size < 2 {
		fmt.Println("Invalid Input")
		return
	}
	var sum float64
	sort.Ints(data)

	minFirst := 0
	minSecond := size - 1

	minSum := math.Abs(float64(data[minFirst] + data[minSecond]))

	for i, j := 0, size-1; i < j; {
		sum = float64(data[i] + data[j])
		if math.Abs(sum) < minSum {
			minSum = math.Abs(sum)
			minFirst = i
			minSecond = j
		}
		if sum < 0 {
			i++
		} else if sum > 0 {
			j--
		} else {
			break
		}
	}
	fmt.Println("The two elements with minimum sum are:", data[minFirst], data[minSecond])
}
