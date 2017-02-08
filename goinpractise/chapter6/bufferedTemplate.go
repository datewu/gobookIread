package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("templates/simple.html"))
}

type page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &page{
		Title:   "An Example",
		Content: "Have fun third .",
	}
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, "A error occured.")
		return
	}
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
