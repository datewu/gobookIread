package main

import (
	"fmt"
	"math/rand"
	"time"
)

func beginningHeartbeat() {
	dowork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		heartbeatStream := make(chan interface{}, 1)
		workStream := make(chan int)

		go func() {
			defer close(heartbeatStream)
			defer close(workStream)

			for i := 0; i < 10; i++ {
				select {
				case heartbeatStream <- struct{}{}:
				default:
				}

				select {
				case <-done:
					return
				case workStream <- rand.Intn(40):
				}

			}
		}()

		return heartbeatStream, workStream
	}
	done := make(chan interface{})
	defer close(done)
	heartbeat, results := dowork(done)

	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Println("results", r)
			} else {
				return
			}
		}
	}

}

func intervalHeartbeat() {
	dowork := func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		results := make(chan time.Time)

		go func() {
			defer close(heartbeat)
			defer close(results)

			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {

				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}

		}()

		return heartbeat, results
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	const timeout = 3 * time.Second
	heartbeat, results := dowork(done, timeout/2)

	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Println("results", r.Second())
		case <-time.After(timeout):
			return
		}

	}
}
