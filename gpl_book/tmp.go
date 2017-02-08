package main

import "fmt"

func main() {
	a := [2]int{1, 3}
	b := [2]int{3, 1}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
}
