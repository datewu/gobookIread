package main

// #include <stdio.h>
// void abc() {
//	printf("Calling C code! the abc() fn\n");
//}
import "C"

import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.abc()
	fmt.Println("Another Go statement!")
}
