package main

import "net/http"

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}

func displayError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "An Error Occurred", http.StatusForbidden)
}
