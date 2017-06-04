package main

import (
	"fmt"
	"os"
	"strings"
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
    An email is  {{ . | lol }}
{{ end }}

{{ with .Jobs}}
  {{ range . }}
    An employer is {{ .Employer }}
	and the role is {{ .Role }}
  {{ end }}
{{ end }}
`

func emailExpander(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}

	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return substrs[0] + " at " + substrs[1]

}

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
	t = t.Funcs(template.FuncMap{"lol": emailExpander})

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
