package main

import (
	"fmt"
	"os"
	"text/template"
)

// Person lol
type Person struct {
	Name   string
	Emails []string
}

const templ = `{"Name": "{{ .Name }}",
"Emails": [
{{ range $index, $elmt := .Emails }}
  {{ if $index }}
    , "{{ $elmt }}"
  {{ else }}
    "{{ $elmt }}"
  {{ end }}
{{ end -}}
]
`

func main() {
	person := Person{
		Name:   "jan",
		Emails: []string{"jan@lol.name", "jan@dota.com"},
	}

	t := template.New("person template")
	t, err := t.Parse(templ)
	checkError(err)

	err = t.Execute(os.Stdout, person)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err)
		os.Exit(3)
	}

}
