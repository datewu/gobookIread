package main

import (
	"fmt"
)

type demo struct{ l int }

func main() {
	var (
		a, b, c = demo{3}, demo{5}, demo{7}
	)
	var (
		s = []*demo{{1}, {3}, {7}}
		p = []*demo{&a, &b, &c}
	)
	fmt.Println(s)
	fmt.Println(p)
}
