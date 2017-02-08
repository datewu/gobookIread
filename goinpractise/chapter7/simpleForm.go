package main

import (
	"fmt"
	"log"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	name := r.FormValue("name")

	fmt.Fprintf(w, name)
}
