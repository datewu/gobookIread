package main

import (
	"fmt"
	"sync"
)

func poolGurad() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			//calcPool.Get()

			// ➜  chapter3 git:(master) ✗ go run pool.go
			// 4 calculators were created.
			// ➜  chapter3 git:(master) ✗ go run pool.go
			// 1017015 calculators were created.

			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Println(numCalcsCreated, "calculators were created.")
}
