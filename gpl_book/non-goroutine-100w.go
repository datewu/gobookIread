package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	m, n := 1000, 1000
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			trivialOPerate(i * j)
		}

	}

}

func trivialOPerate(non int) {
	fmt.Fprintf(ioutil.Discard, "dump dump %d", non)

}
