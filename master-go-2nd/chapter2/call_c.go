package main

// #cgo CFLAGS: -I${SRCDIR}/libC
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(("Going to call a C func"))
	C.cHello()

	fmt.Println(("Going to call another C func with args"))
	msg := C.CString("Lo gin")
	defer C.free(unsafe.Pointer(msg))
	C.printMsg(msg)

	fmt.Println("Done")
}

// ➜  chapter2 git:(master) ✗ gcc -c libC/*.c
// ➜  chapter2 git:(master) ✗ ar rs callC.a callC.o
// ➜  chapter2 git:(master) ✗ go run call_c.go
