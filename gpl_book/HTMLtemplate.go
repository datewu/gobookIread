package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const issuesURL = "https://api.github.com/search/issues"

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

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
	if err := issueList.Execute(os.Stdout, r); err != nil {
		log.Fatalln(err)
	}
}

// go run HTMLtemplate.go react redux angular js  > test.html
