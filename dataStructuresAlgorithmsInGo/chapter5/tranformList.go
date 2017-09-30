package main

// “How would you swap elements of a list like [a1 a2 a3 a4 b1 b2 b3 b4] to convert it into [a1 b1 a2 b2 a3 b3 a4 b4]?
// Approach:
// ·         First swap elements in the middle pair
// ·         Next swap elements in the middle two pairs
// ·         Next swap elements in the middle three pairs
// ·         Iterate n-1 steps.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

func transform(str string) string {
	data := []rune(str)
	size := len(data)
	n := size / 2
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			data[n-i+2*j], data[n-i+2*j+1] = data[n-i+2*j+1], data[n-i+2*j]
		}
	}
	return string(data)
}
