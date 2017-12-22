package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)

	} else {
		for _, fname := range files {
			f, err := os.Open(fname)
			if err != nil {
				log.Println("dup:", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}

	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, c map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		c[input.Text()]++
	}

}
