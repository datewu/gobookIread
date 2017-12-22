package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	sizeStream := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, sizeStream)
		}
		close(sizeStream)
	}()

	var nfiles, nbytes int64
	for size := range sizeStream {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)

}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1fGB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, fileSizeStream chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizeStream)
		} else {
			fileSizeStream <- entry.Size()
		}

	}

}
