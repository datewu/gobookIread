package main

//
// “Given a 2 dimensional list. Each row and column are sorted in ascending order. How would you find an element in it?
//
// The algorithm works as:
// 1.    Start with element at last column and first row
// 2.    If the element is the value we are looking for, return true.
// 3.    If the element is greater than the value we are looking for, go to the element at previous column but same row.
// 4.    If the element is less than the value we are looking for, go to the element at next row but same column.
// 5.    Return false, if the element is not found after reaching the element of the last row of the first column. Condition (row < r && column >= 0) is false.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.

func findIn2D(data [][]int, r, c, value int) bool {
	row, column := 0, c-1
	for row < r && column >= 0 {
		if value == data[row][column] {
			return true
		} else if value < data[row][column] {
			column--
		} else {
			row++
		}
	}
	return false
}
