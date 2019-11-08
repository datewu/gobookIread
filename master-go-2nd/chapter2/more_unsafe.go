package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [...]int{0, 1, -2, 3, 4}
	p := &a[0]
	fmt.Print(*p, " ")
	mAddr := uintptr(unsafe.Pointer(p)) +
		unsafe.Sizeof(a[0])
	for index := 0; index < len(a)-1; index++ {
		p = (*int)(unsafe.Pointer(mAddr))
		fmt.Print(*p, " ")
		mAddr = uintptr(unsafe.Pointer(p)) +
			unsafe.Sizeof(a[0])
	}
	fmt.Println()
	p = (*int)(unsafe.Pointer(mAddr))
	fmt.Print("One more: ", *p, " ")
	mAddr = uintptr(unsafe.Pointer(p)) +
		unsafe.Sizeof(a[0])
	fmt.Println("same???", *p)

}
