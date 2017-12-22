package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) {
	var root *tree
	//root := new(tree)
	for _, v := range values {
		root = add(root, v)
	}
	// fmt.Println(appendValues(values[:0], root))
}

func appendValues(vs []int, t *tree) (res []int) {
	if t != nil {
		vs = appendValues(vs, t.left)
		vs = append(vs, t.value)
		vs = appendValues(vs, t.right)
	}
	return vs
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	randValues := rand.Perm(8)
	fmt.Println("Original", randValues)

	sort(randValues)
	fmt.Println("Sorted", randValues)
}
