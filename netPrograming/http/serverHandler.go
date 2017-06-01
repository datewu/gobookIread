package main

import "net/http"

func main() {
	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	http.ListenAndServe(":8080", myHandler)
}
