package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("templates/simple.html"))

type page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := page{
		Title:   "An example",
		Content: "Have fun again",
	}
	t.Execute(w, p)
}
func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
