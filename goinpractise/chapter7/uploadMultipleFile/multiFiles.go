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
		t, _ := template.ParseFiles("files.html")
		t.Execute(w, nil)
	} else {
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			log.Println(err)
			fmt.Fprintln(w, err)
			return
		}
		data := r.MultipartForm
		files := data.File["file"]
		for _, fHeader := range files {
			f, err := fHeader.Open()
			if err != nil {
				log.Println(err)
				fmt.Fprintln(w, err)
				return
			}
			filename := fHeader.Filename + ".multi.upload"
			out, err := os.Create(filename)
			if err != nil {
				log.Println(err)
				fmt.Fprintln(w, err)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, f)
			if err != nil {
				log.Println(err)
				fmt.Fprintln(w, err)
				return
			}
		}
		fmt.Fprintf(w, "Upoad complate")
	}
}
