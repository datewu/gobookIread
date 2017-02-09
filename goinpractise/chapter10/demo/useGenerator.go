//go:generate ./queue myInt
package main

import "fmt"

type myInt int

func main() {
	var one, two, three myInt = 1, 2, 5
	q := newmyIntQueue()
	q.Insert(one)
	q.Insert(two)
	q.Insert(three)

	fmt.Println(q)
	fmt.Printf("First value: %d\n", q.Remove())
	fmt.Println(q)
}
