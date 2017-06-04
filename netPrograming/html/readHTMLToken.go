package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usgae:", os.Args[0], "file")
	}

	file := os.Args[1]
	bytes, err := ioutil.ReadFile(file)
	checkError(err)

	r := strings.NewReader(string(bytes))
	z := html.NewTokenizer(r)

	var depth int
	for {
		tt := z.Next()
		for n := 0; n < depth; n++ {
			fmt.Print(" ")
		}
		switch tt {
		case html.ErrorToken:
			fmt.Println("Error", z.Err())
			os.Exit(0)
		case html.TextToken:
			fmt.Println("Text: \"" + z.Token().String() + "\"")
		case html.StartTagToken, html.EndTagToken:
			fmt.Println("Tag: \"" + z.Token().String() + "\"")
			if tt == html.StartTagToken {
				depth++
			} else {
				depth--
			}
		}

	}

}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err)
	}

}
