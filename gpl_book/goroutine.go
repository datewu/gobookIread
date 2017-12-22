package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	m, n := 100, 100

	wg.Add(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			go func(z int) {
				trivialOPerate(z)
				wg.Done()
			}(i * j)
		}

	}
	wg.Wait()

}

func trivialOPerate(non int) {
	time.Sleep(time.Millisecond)
	fmt.Fprintf(ioutil.Discard, "dump dump %d", non)

}
