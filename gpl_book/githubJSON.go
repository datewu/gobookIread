package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const issuesURL = "https://api.github.com/search/issues"

type issuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*issue
}

type issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *user
	CreateAt time.Time `json:"create_at"`
}

type user struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func searchIssues(terms []string) (*issuesSearchResult, error) {
	qraw := strings.Join(terms, " ")
	q := url.QueryEscape(qraw)
	resp, err := http.Get(issuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query: %s failed: %s", qraw, resp.Status)
	}
	var result issuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	r, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d isses:\n", r.TotalCount)
	for _, item := range r.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

// go run githibJSON.go react redux angular js
