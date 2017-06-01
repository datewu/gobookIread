package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host:port")
		os.Exit(3)
	}
	url := os.Args[1]
	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k, ":", v)
	}
}
