package main

import (
	"errors"
	"io"
	"log"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := openCSV("data.csv")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
}

func openCSV(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	file, err = os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	removeEmptyLines(file)
	return
}

func removeEmptyLines(f *os.File) {
	panic(errors.New("Failed parse"))
}
