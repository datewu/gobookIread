package main

import (
	"fmt"
	"log"
)

func permutation(data []int, i, length int) {
	if length == i {
		fmt.Println(data)
		return
	}

	for j := i; j < length; j++ {
		swap(data, i, j)
		permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func swap(data []int, x, y int) {
	if x != y {
		data[x], data[y] = data[y], data[x]
	}
}

func main() {
	d := []int{1, 2, 3, 10, 9, 6}
	fmt.Println(d)
	permutation(d, 0, len(d))
	log.Println(3)
	permutation(d, 0, 3)
}
