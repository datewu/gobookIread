package main

import "C"

import (
	"fmt"
)

//export Printmsg
func Printmsg() {
	fmt.Println("A Go function")
}

//export Mul
func Mul(a, b int) int {
	return a * b
}

func main() {
}

// must use
//export Name
// or
// there is no usedByC.h file generated

// go build -o usedByC.o -buildmode=c-shared usedByC.go
