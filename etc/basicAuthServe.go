package main

import (
	"crypto/subtle"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", basicAuth(hello, "c", "d", "password due"))
	http.ListenAndServe(":9999", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<div id="lol">"due, what's up?"</div>`)
}

func basicAuth(h http.HandlerFunc, username, password, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised. Due\n"))
			return
		}
		h(w, r)
	}
}
