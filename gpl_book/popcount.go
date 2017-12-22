package main

import "fmt"

var pCount [256]byte

func main() {
	var i uint64 = 5
	for ; i < 1<<62; i <<= 4 { // chang to 1<<63 you will get infinite loop :)
		fmt.Println(i, popCount(i))
	}
}

func popCount(x uint64) int {
	var b byte
	for i := 0; i < 8; i++ {
		index := byte(x >> (8 * uint(i)))
		b += pCount[index]
	}
	return int(b)
}

func init() {
	for i := range pCount {
		// pCount[i] is the population count of i.
		pCount[i] = pCount[i/2] + byte(i&1)
	}

}
