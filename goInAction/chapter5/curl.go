// Sample program to show how to write a simple version of
// curl using io.Reader and io.Writer interface support.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./crul <url>")
		os.Exit(-1)
	}
}

func main() {
	// Get a response from the web server
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// Copies from the Body to Stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
