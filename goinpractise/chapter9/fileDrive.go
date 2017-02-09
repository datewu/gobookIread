package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// File LOL
type File interface {
	Load(string) (io.ReadCloser, error)
	Save(string, io.ReadSeeker) error
}

type localFile struct {
	Base string
}

func (l localFile) Load(path string) (io.ReadCloser, error) {
	p := filepath.Join(l.Base, path)
	return os.Open(p)
}

func (l localFile) Save(path string, body io.ReadSeeker) error {
	p := filepath.Join(l.Base, path)
	d := filepath.Dir(p)
	err := os.MkdirAll(d, os.ModeDir|os.ModePerm)
	if err != nil {
		log.Println(err)
		return err
	}
	f, err := os.Create(p)
	if err != nil {
		log.Println(err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, body)
	return err
}

func main() {
	content := `date for test and seed "" $%&^(*&)(@$`
	body := bytes.NewReader([]byte(content))

	store, err := fileStore()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Storing content...")
	err = store.Save("foo/bar", body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Retrieving  content...")
	c, err := store.Load("foo/bar")
	if err != nil {
		log.Println(err)
		return
	}
	o, err := ioutil.ReadAll(c)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(o))
}

func fileStore() (File, error) {
	return &localFile{Base: "."}, nil
}
