package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", lll)
	http.ListenAndServe(":8080", nil)
}

func lll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
