package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver()
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	http.ListenAndServe(":8080", pr)
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range p.handlers {
		ok, err := path.Match(pattern, check)
		if ok && err == nil {
			handlerFunc(w, r)
			return
		}
		if err != nil {

			fmt.Fprint(w, err)
			return
		}
	}
	http.NotFound(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "lol a great game"
	}
	fmt.Fprint(w, "hello ", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "dota also a great game"
	}
	fmt.Fprint(w, "Goodbye", name)
}
