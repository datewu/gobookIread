package main

//
// “Bucket sort is the simplest and most efficient type of sorting. Bucket sort has a strict requirement of a predefined range of data.
// Like, sort how many people are in which age group. We know that the age of people can vary between 0 and 130.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func bucketSort(data []int, lRange, uRange int) {
	rng := uRange - lRange
	size := len(data)
	count := make([]int, rng)
	for i := 0; i < size; i++ {
		count[data[i]-lRange]++
	}
	for i, j := 0, 0; i < rng; i++ {
		for ; count[i] > 0; count[i]-- {
			data[j] = i + lRange
			j++
		}
	}
}
