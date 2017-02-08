package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * time.Duration(i))
	}
}

func main() {
	go count()
	time.Sleep(time.Millisecond * 5)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 20)
	fmt.Println("Ending World")
}
