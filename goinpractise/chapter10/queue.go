package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl = `
package {{ .Package }}

type {{ .MyType }}Queue struct {
	q []{{ .MyType }}
}

func new{{ .MyType }}Queue() *{{ .MyType }}Queue {
	return &{{ .MyType }}Queue{
		q: []{{ .MyType }}{},
	}
}

func (q *{{ .MyType }}Queue) Insert(v {{ .MyType }}) {
	q.q = append(q.q, v)
}

func (q *{{ .MyType }}Queue) Remove()(res {{ .MyType }}) {
	if len(q.q) == 0 {
		panic("Oops")
	}
	res = q.q[0]
	q.q = q.q[1:]
	return
}
`

func main() {
	tt := template.Must(template.New("queue").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)
		if err != nil {
			log.Println(err)
			continue
		}
		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		tt.Execute(file, vals)
		file.Close()
	}
}
