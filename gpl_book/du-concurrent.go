package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messsage")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	wg.Add(len(roots))

	sizeStream := make(chan int64)
	go func() {
		for _, root := range roots {
			go walkDir(root, &wg, sizeStream)
		}
	}()

	go func() {
		wg.Wait()
		close(sizeStream)
	}()

	var nfiles, nbytes int64
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-sizeStream:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}

	}
	printDiskUsage(nfiles, nbytes)

}

var semaStream = make(chan struct{}, 80)

func dirents(dir string) []os.FileInfo {
	semaStream <- struct{}{}
	defer func() {
		<-semaStream
	}()
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

func walkDir(dir string, n *sync.WaitGroup, fileSizeStream chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizeStream)
		} else {
			fileSizeStream <- entry.Size()
		}

	}

}
