package main

import "fmt"

func main() {
	fmt.Println(s)
	fmt.Println(p)
}

type lol struct {
	l int
}

var s = []*lol{{1}, {2}, {5}}

var p = []*lol{&a, &b, &c}

var (
	a = lol{4}
	b = lol{9}
	c = lol{2}
)
