package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var elements []string

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			elements = append(elements, tok.Name.Local) // push
		case xml.EndElement:
			elements = elements[:len(elements)-1]
		case xml.CharData:
			if containsAll(elements, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(elements, " "), tok)
			}
		}

	}

}
