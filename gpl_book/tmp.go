package main

import "fmt"

func main() {
	a := [2]int{1, 5}
	b := [2]int{5, 1}
	c := [2]int{1, 5}
	fmt.Println(a == b, a == c, b == c)
	fmt.Println(a, b, c)
	fmt.Printf("addresses: %p, %p, %p\n", a, b, c)
	fmt.Printf("addresses: %p, %p, %p", &a, &b, &c)
}
