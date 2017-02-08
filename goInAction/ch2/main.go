package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	_ "./matchers"
	"./search"
)

func main() {
	// Perform the search for the specified term.
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("input key words, Please")
	for s.Scan() {
		search.Run(s.Text())
	}
}

func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}
