package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("outside  a goroutine.")
	go func() {
		fmt.Println("Inside a goroutine.")
	}()
	fmt.Println("outside  again.")

	runtime.Gosched()

}
