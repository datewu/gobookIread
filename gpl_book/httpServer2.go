package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

var (
	mu    sync.Mutex
	count int
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
