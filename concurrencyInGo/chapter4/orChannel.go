package main

import (
	"fmt"
	"time"
)

func orChannel() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	<-orChan(
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
		sig(7*time.Second),
		sig(6*time.Second),
		sig(2*time.Second),
	)

	fmt.Println("done after", time.Since(start))

}

var orChan func(channels ...<-chan interface{}) <-chan interface{}

func init() {
	orChan = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		orDone := make(chan interface{})

		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-orChan(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
}
