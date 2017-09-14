package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lexicalConfinement() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i

			}
		}()
		return results
	}

	consumer := func(r <-chan int) {
		for result := range r {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	upStream := chanOwner()

	consumer(upStream)
}

func goroutineLeaks() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)

	fmt.Println("done.")

}

func eliminatedGoroutineLeaks() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return

				}
			}
		}()
		return completed
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("done.")
}

func producerCannel() {
	randStream := func(done <-chan interface{}) <-chan int {
		rStream := make(chan int)
		go func() {
			defer fmt.Println("randstream closure exited.")
			close(rStream)
			for {
				select {
				case rStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return rStream
	}

	done := make(chan interface{})
	producer := randStream(done)
	fmt.Println("3 random  ints:")

	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-producer)
	}
	close(done)
	time.Sleep(5 * time.Second)
}
