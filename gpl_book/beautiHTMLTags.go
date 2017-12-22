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
	forEachNode(beforeElement, afterElement, doc)
}

func forEachNode(pre, post func(*html.Node), n *html.Node) {

	if pre != nil {
		pre(n)
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(pre, post, child)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func beforeElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func afterElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
