package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var v int64 = 5
	var p1 = &v
	var p2 = (*int32)(unsafe.Pointer(p1))
	fmt.Println("*p1:", *p1)
	fmt.Println("*p2:", *p2)
	*p1 = 953452342348787913
	fmt.Println((v))
	fmt.Println("*p2:", *p2)

	*p1 = 998
	fmt.Println((v))
	fmt.Println("*p2:", *p2)
}
