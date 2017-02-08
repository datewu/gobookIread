package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, my name is noone")
}
func main() {
	http.HandleFunc("/", hello)
	log.Println(http.ListenAndServe(":4000", nil))
}
