package main

import (
	"fmt"
	"time"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i < 7; i++ {
				resultStream <- i

			}
		}()
		return resultStream
	}

	r := chanOwner()
	for result := range r {
		fmt.Println("Received:", result)
	}
	fmt.Println("Done reveiving")

	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()
	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCounter++
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Achieved", workCounter, "cycles of work before signalled to stop")
}
