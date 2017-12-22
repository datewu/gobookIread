package main

import (
	"fmt"
	"time"
)

func main() {
	pipleLine3()
}

func pipleLine1() {
	naturalsStream := make(chan int)
	squaresStream := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturalsStream <- x
		}
	}()

	go func() {
		for {
			x := <-naturalsStream
			squaresStream <- x * x
		}

	}()

	for {
		fmt.Println(<-squaresStream)
		time.Sleep(time.Second)
	}

}

func pipleLine2() {
	naturalsStream := make(chan int)
	squaresStream := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturalsStream <- x
		}
		close(naturalsStream)
	}()

	go func() {
		for x := range naturalsStream {
			squaresStream <- x * x

		}
		close(squaresStream)
	}()

	for n := range squaresStream {
		time.Sleep(time.Second)
		fmt.Println(n)
	}
}

func pipleLine3() {
	naturalsStream := make(chan int)
	squaresStream := make(chan int)

	counter := func(outStream chan<- int) {
		for x := 0; x < 100; x++ {
			outStream <- x
		}
		close(outStream)
	}

	squarer := func(outStream chan<- int, inStream <-chan int) {
		for v := range inStream {
			outStream <- v * v
		}
		close(outStream)
	}

	printer := func(inStream <-chan int) {
		for v := range inStream {
			fmt.Println(v)
		}
	}
	go counter(naturalsStream)
	go squarer(squaresStream, naturalsStream)
	printer(squaresStream)

}
