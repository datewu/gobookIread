package main

import "fmt"

func main() {
	n1, n3 := Names()
	fmt.Println(n1, n3)
}

// Names lol
func Names() (string, string) {
	return "Foo", "Bar"
}
