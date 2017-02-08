package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var (
	cache map[string]*cacheFile
	mutex = new(sync.RWMutex)
)

func main() {
	cache = make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	mutex.RLock()
	v, found := cache[key]
	mutex.RUnlock()
	if !found {
		mutex.Lock()
		defer mutex.Unlock()
		fileName := "./files" + key
		f, err := os.Open(fileName)
		defer f.Close()
		if err != nil {
			log.Println(err)
			http.NotFound(w, r)
			return
		}

		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			log.Println(err)
			http.NotFound(w, r)
			return
		}
		r := bytes.NewReader(b.Bytes())
		info, _ := f.Stat()
		v = &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[key] = v
	}
	http.ServeContent(w, r, key, v.modTime, v.content)
}
