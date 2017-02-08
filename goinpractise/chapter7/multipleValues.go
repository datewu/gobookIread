package main

import (
	"fmt"
	"log"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	maxMemory := 16 << 20
	err := r.ParseMultipartForm(maxMemory)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	for k, v := range r.PostForm["name"] {
		log.Println(k, v)
	}

	fmt.Fprintf(w, r.PostForm["name"])
}
