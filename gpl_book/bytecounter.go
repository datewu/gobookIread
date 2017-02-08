package main

import "fmt"

type byteCounter int

func main() {
	var c byteCounter
	fmt.Println(c)
	c.Write([]byte("hello"))
	fmt.Println(c)
	name := "dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

func (c *byteCounter) Write(p []byte) (int, error) {
	*c += byteCounter(len(p))
	//return 2, nil
	return len(p), nil
}
