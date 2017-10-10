package main

func maxSubArraySum(data []int) (sum int) {
	size := len(data)

	maxEndingHere := 0
	for i := 0; i < size; i++ {
		maxEndingHere += data[i]
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
		if sum < maxEndingHere {
			sum = maxEndingHere
		}
	}
	return
}

// func main() {
// 	d := []int{1, 4, -10, 998, 997, -10000, 3, 5, -6, 7, 9, 9, -3}
//
// 	fmt.Println(d)
// 	fmt.Println(maxSubArraySum(d))
// }
