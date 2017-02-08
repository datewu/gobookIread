package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	timeout := time.After(5 * time.Second)

	go send(ch)

	for {
		select {
		case m := <-ch:
			if m {
				fmt.Println("Got message")
			} else {
				return
			}

		case <-timeout:
			fmt.Println("Time out")
			return
		default:
			println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func send(ch chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(820 * time.Millisecond)
		ch <- true
	}
	close(ch)
	println("Sent and closed")
}
