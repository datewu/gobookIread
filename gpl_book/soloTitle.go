package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type bailout struct{}

// solo returns the text of the first non-empty tile element
// in doc, and an error if these is not exactly one.
func sole(url string) (title string, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			err = errors.New("multiple title elements")
		default:
			panic(p)

		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		err = fmt.Errorf("%s has type %s, not text/html", url, ct)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data

		}
	}
	forEachNode(doc, visitNode, nil)
	return
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	u := "http://jd.com"
	fmt.Println(sole(u))
}
