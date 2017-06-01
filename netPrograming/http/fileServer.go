package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// deliver files from the directory /var/www
	fileServer := http.FileServer(http.Dir("/var/www"))
	// register the handler and deliver requests to it
	err := http.ListenAndServe(":8080", fileServer)
	fmt.Println("lol")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err)
		os.Exit(3)
	}
}
