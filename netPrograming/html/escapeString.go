package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", escapeString)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func escapeString(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	bytes, err := ioutil.ReadFile("." + r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	escapedStr := html.EscapeString(string(bytes))
	htmlText := "<html><body><pre><code>" +
		escapedStr +
		"</code></pre></body></html>"
	w.Write([]byte(htmlText))
}
