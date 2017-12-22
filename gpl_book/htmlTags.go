package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalln("HTML parsing err:", err)
	}
	outline(nil, doc)
}

func outline(tags []string, n *html.Node) {
	if n.Type == html.ElementNode {
		tags = append(tags, n.Data)
		fmt.Println(tags)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(tags, c)
	}
}
