package main

import "fmt"

type tree struct {
	LeafValue int
	Right     *tree
	Left      *tree
}

func main() {
	root := tree{
		LeafValue: 9,
		Right: &tree{
			LeafValue: 4,
			Right:     &tree{4, nil, nil},
			Left:      nil,
		},
		Left: &tree{7, nil, nil},
	}

	fmt.Println(root.Right.Right.LeafValue)
}
