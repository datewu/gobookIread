package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print("d1:", i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print("d2:", i, " ")
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print("d3:", n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()

	d3()
	fmt.Println()
}

// ➜  chapter2 git:(master) ✗ go run defer.go
// d1:1 d1:2 d1:3 d2:0 d2:0 d2:0
// d3:1 d3:2 d3:3
