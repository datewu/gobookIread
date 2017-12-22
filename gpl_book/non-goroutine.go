package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	m, n := 100, 100
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			trivialOPerate(i * j)
		}

	}

}

func trivialOPerate(non int) {
	time.Sleep(time.Millisecond)
	fmt.Fprintf(ioutil.Discard, "dump dump %d", non)

}
