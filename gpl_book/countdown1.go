package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown, <-tick)
	}
	fmt.Printf("lol")
}
