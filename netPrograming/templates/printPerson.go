package main

import (
	"fmt"
	"os"
	"text/template"
)

// Person lol
type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
}

// Job lol
type Job struct {
	Employer string
	Role     string
}

const templ = `The name is {{ .Name }}.
The age is {{ .Age }}.
{{ range .Emails }}
    An email is  {{ . }}
{{ end }}

{{ with .Jobs}}
  {{ range . }}
    An employer is {{ .Employer }}
	and the role is {{ .Role }}
  {{ end }}
{{ end }}
`

func main() {
	job1 := Job{Employer: "Box Hill Institure", Role: "Director, Commerce and ICT"}
	job2 := Job{Employer: "Canberra University", Role: "Agjunct Professor"}

	person := Person{
		Name:   "jan",
		Age:    99,
		Emails: []string{"jan@lol.name", "jan@dota.com"},
		Jobs:   []*Job{&job1, &job2},
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
