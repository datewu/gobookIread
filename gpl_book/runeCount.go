package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "runeCount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune \t count\n")
	for r, n := range counts {
		fmt.Printf("%q\t%d\n", r, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for length, count := range utflen {
		if length > 0 {
			fmt.Printf("%d\t%d\n", length, count)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n %d invalid UTF-8 characters \n", invalid)
	}
}
