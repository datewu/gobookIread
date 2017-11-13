package main

func rotateArray(data []int, k int) {
	n := len(data)
	reverseArray(data, 0, k-1)
	reverseArray(data, k, n-1)
	reverseArray(data, 0, n-1)
}

func reverseArray(data []int, start, end int) {
	i, j := start, end
	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}

//
// func main() {
// 	s := []int{1, 5, 76, 8, 9, 4, 9, 8, 3, 5, 67, 4}
// 	fmt.Println(s)
// 	rotateArray(s, 4)
// 	fmt.Println(s)
// }
