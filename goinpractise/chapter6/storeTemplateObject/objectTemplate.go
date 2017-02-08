package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type page struct {
	Title   string
	Content template.HTML
}

type quote struct {
	Quote, Person string
}

func main() {
	q := &quote{
		`You keep using that word. I do not think
		it means what you think it means.`,
		"Inigo Montoya",
	}
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String())

	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &page{
		"Some User",
		qc,
	}
	t.ExecuteTemplate(w, "index.html", p)
}
