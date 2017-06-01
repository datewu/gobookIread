package main

import "net/http"

func main() {
	// deliver files from the directory /var/www
	fileServer := http.FileServer(http.Dir("/var/www"))
	// register the handler and deliver requests to it
	http.ListenAndServeTLS(":8080", "domain.pem", "private.pem", fileServer)
}
