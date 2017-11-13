package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func stackDemo() {
	s := stack.New()
	s.Push(4)
	s.Push(5)
	s.Push(7)
	s.Push(8)
	s.Push(9)

	for s.Len() != 0 {
		fmt.Print(s.Pop())
	}

}
