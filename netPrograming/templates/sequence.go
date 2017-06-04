package main

import (
	"errors"
	"fmt"
	"os"
	"text/template"
)

var fmp = template.FuncMap{
	"sequence": sequenceFunc,
	"cycle":    cycleFunc,
}

var tmpl = `{{ $comma := sequence "" ", " }}
{{ range $ }} {{ $comma.Next }} {{ . }} {{ end }}
{{ $colour := cycle "black" "white" "red" }}
{{ range $ }} {{ $comma.Next }} {{ . }} in {{ $colour.Next }} {{ end }}
`

func main() {
	t, err := template.New("").Funcs(fmp).Parse(tmpl)
	if err != nil {
		fmt.Printf("parse error: %v\n", err)
		return
	}
	err = t.Execute(os.Stdout, []string{"a", "b", "x", "y", "z"})
	if err != nil {
		fmt.Printf("parse error: %v\n", err)
	}
}

type generator struct {
	ss []string
	i  int
	f  func(s []string, i int) string
}

func (seq *generator) Next() string {
	s := seq.f(seq.ss, seq.i)
	seq.i++
	return s
}

func sequenceGen(ss []string, i int) string {
	if i >= len(ss) {
		return ss[len(ss)-1]
	}
	return ss[i]
}

func sequenceFunc(ss ...string) (*generator, error) {
	if len(ss) == 0 {
		return nil, errors.New("sequence must have at least one element")
	}
	return &generator{ss, 0, sequenceGen}, nil
}

func cycleGen(ss []string, i int) string {
	return ss[i%len(ss)]
}

func cycleFunc(ss ...string) (*generator, error) {
	if len(ss) == 0 {
		return nil, errors.New("sequence must have at least one element")
	}
	return &generator{ss, 0, cycleGen}, nil
}
