package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	res := sum(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, res)
}

func sum(a, b int) int {
	return a + b
}
func multiply(a, b int) int {
	return a * b
}
