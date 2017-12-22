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

const templ = `
{{ .TotalCount }} issues:
{{ range .Items }} ------------------
Number: {{ .Number }}
User: {{ .User.Login }}
Title: {{ .Title | printf "%.64s" }}
Age: {{ .CreateAt | daysAgo }} days
{{ end }}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

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
	CreateAt time.Time `json:"created_at"`
	Body     string    // markdown format
}

type user struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func searchIssues(terms []string) (r *issuesSearchResult, err error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issuesURL + "?q=" + q)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("search query failed %s", resp.Status)
		return
	}
	// var res issuesSearchResult
	// if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
	// 	return
	// }
	// r = &res
	m := new(issuesSearchResult)
	if err = json.NewDecoder(resp.Body).Decode(m); err != nil {
		return
	}
	r = m
	return
}
func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}
	if err = report.Execute(os.Stdout, result); err != nil {
		log.Fatalln(err)
	}

}
