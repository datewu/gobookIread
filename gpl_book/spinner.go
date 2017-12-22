package main

import (
	"fmt"
	"time"
)

var doneStream = make(chan struct{})

func main() {
	go spinner(100 * time.Microsecond)
	n := 45
	fibN := fib(n)
	<-doneStream
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

}

func spinner(delay time.Duration) {
	for {
		if done() {
			return
		}
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)

}

func done() bool {
	select {
	case doneStream <- struct{}{}:
		return true
	default:
		return false
	}
}
