package main

import (
	"fmt"
	"sync"
)

// Button broadcast cond
type Button struct {
	Clicked *sync.Cond
}

func bcast() {
	button := Button{sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegisterd sync.WaitGroup
	clickRegisterd.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Display annoying dialog box")
		clickRegisterd.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Display annoying dialog box again")
		clickRegisterd.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Display annoying dialog box the third time")
		clickRegisterd.Done()
	})
	button.Clicked.Broadcast()
	clickRegisterd.Wait()

}
