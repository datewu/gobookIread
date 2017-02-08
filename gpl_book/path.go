package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := path.Base(r.URL.Path)
		log.Println(strings.ToLower(s))
		fmt.Fprintf(w, "%s\n", s)
	})
	http.ListenAndServe(":8080", nil)
}
