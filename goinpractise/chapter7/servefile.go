package main

import "net/http"

func main() {
	http.HandleFunc("/", readme)
	http.ListenAndServe(":8080", nil)
}

func readme(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/readme.txt")
}
