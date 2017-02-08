package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

type page struct {
	Title, Content string
}

func init() {
	t = template.Must(template.ParseFiles("index.html", "head.html"))
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &page{
		"An Example LOL",
		"Have fun bike",
	}
	t.ExecuteTemplate(w, "index.html", p)
}
func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
