package main

import "fmt"

func names() (first, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := names()
	fmt.Println(n1, n2)
}
