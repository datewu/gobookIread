package main

//
// “Given a list of even and odd numbers, write a program to separate even numbers from the odd numbers.
//
// First approach: allocate a separate list, then scan through the given list, and fill even numbers from the start and odd numbers from the end.
//
// Second approach: Algorithm is as follows.
// 1.    Initialize the two variable left and right. Variable left=0 and right= size-1.
// 2.    Keep increasing the left index until the element at that index is even.
// 3.    Keep decreasing the right index until the element at that index is odd.
// 4.    Swap the number at left and right index.
// 5.    Repeat steps 2 to 4 until left is less than right.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func seperate(data []int) {
	size := len(data)
	left, right := 0, size-1
	for left < right {
		if data[left]%2 == 0 {
			left++
		} else if data[right]%2 == 1 {
			right--
		} else {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}
	}
}
