package main

import (
	"html/template"
	"net/http"
)

var t map[string]*template.Template

func init() {
	t = make(map[string]*template.Template)
	temp := template.Must(template.ParseFiles("base.html", "user.html"))
	t["user.html"] = temp

	temp = template.Must(template.ParseFiles("base.html", "page.html"))
	t["page.html"] = temp
}

type page struct {
	Title, Content string
}

type user struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &page{
		"An Example",
		"Have page",
	}
	t["page.html"].ExecuteTemplate(w, "base", p)
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &user{
		"swordsmith",
		"lol dashen",
	}
	t["user.html"].ExecuteTemplate(w, "base", u)
}
func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}
