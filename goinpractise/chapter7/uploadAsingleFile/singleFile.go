package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", fileForm)
	http.ListenAndServe(":8080", nil)
}

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.html")
		t.Execute(w, nil)
	} else {
		f, header, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		filename := header.Filename + ".untrust.upload"
		out, err := os.Create(filename)
		if err != nil {
			log.Println(err)
			http.NotFound(w, r)
			return
		}
		defer out.Close()

		io.Copy(out, f)
		fmt.Fprintf(w, "Upoad complate")
	}
}
