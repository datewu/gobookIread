package main

import "fmt"

func tohUtil(num int, from, to, temp string) {

	if num < 1 {
		return
	}
	tohUtil(num-1, from, temp, to)
	fmt.Println("Move disk", num, "from peg", from, "to peg", to)
	tohUtil(num-1, temp, to, from)
}

// func main() {
// 	fmt.Println("The sequence of moves involved in the Tower of Hanoi are:")
// 	tohUtil(10, "A", "C", "B") // Need 2^10 - 1 = 1023 steps
// }
