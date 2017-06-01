package main

import (
	"net/http"
	"os"
)

func main() {
	// file handler for most files
	fileServer := http.FileServer(http.Dir("/var/www"))
	http.Handle("/", fileServer)

	// function handler for /cgi-bin/printenv
	http.HandleFunc("/cgi-bin/printenv", printenv)

	// deliver request to the handlers
	http.ListenAndServe(":8080", nil)
}

func printenv(w http.ResponseWriter, r *http.Request) {
	env := os.Environ()
	w.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		w.Write([]byte(v + "\n"))
	}
	w.Write([]byte("</pre>"))
}
