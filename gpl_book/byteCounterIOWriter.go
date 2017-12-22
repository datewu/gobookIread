package main

import (
	"fmt"
)

type byteCounter int

func main() {
	var c byteCounter
	fmt.Println(c)
	c.Write([]byte("hello world"))
	fmt.Println(c)

	name := "dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	c.Write([]byte("done"))
	fmt.Println(c)
}

// Write implement io.Writer interface
func (c *byteCounter) Write(p []byte) (int, error) {
	*c += byteCounter(len(p))
	return len(p), nil
}
