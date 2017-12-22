package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {

}

func fetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatalln("fetch: reading:", url, err)
		}
		fmt.Printf("%s", b)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return
	}
	n, err = io.Copy(f, resp.Body)

	if closeErr := f.Close(); closeErr != nil {
		err = closeErr
	}
	return
}

func fetchall() {
	start := time.Now()
	stream := make(chan string)
	for _, url := range os.Args[1:] {
		go func(u string, s chan<- string) {
			ls := time.Now()
			resp, err := http.Get(u)
			if err != nil {
				s <- err.Error()
				return
			}
			nbytes, err := io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			if err != nil {
				s <- fmt.Sprintf("while reading %s: %v", u, err)
				return
			}
			secs := time.Since(ls).Seconds()
			s <- fmt.Sprintf("%.fs %7d %s", secs, nbytes, u)

		}(url, stream)
	}

	for range os.Args[1:] {
		fmt.Println(<-stream)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
