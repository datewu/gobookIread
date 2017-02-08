package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	tick := time.Tick(time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		case <-tick:
			fmt.Println(countdown)
		}
	}
	fmt.Printf("lol")
}
