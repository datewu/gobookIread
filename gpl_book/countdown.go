package main

import "os"
import "fmt"
import "time"

func main() {
	abortStream := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abortStream <- struct{}{}
	}()

	fmt.Println("Comming countdown....")
	tick := time.Tick(time.Second)
	for i := 0; i < 10; i++ {
		select {
		case <-abortStream:
			fmt.Println("Launch aborted!")
			return
		case <-tick:
			fmt.Println(10 - i)
		}
	}
	fmt.Println("LOL")
}
