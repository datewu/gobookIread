package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("outside  1 goroutine.")
	go func() {
		runtime.Gosched()
		fmt.Println("Inside 1 goroutine.")
	}()

	runtime.Gosched()
	fmt.Println("outside  2 again.")

	runtime.Gosched()
	fmt.Println("outside  3 again.")

}
