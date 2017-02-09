package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name string
}

var stringJSON = `{
	"name": "Miracle Max"
}`

func main() {
	var p person
	err := json.Unmarshal([]byte(stringJSON), &p)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
