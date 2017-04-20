package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p := "packt"
	s := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{"lol", 98}
	jsonP, err := json.Marshal(p)
	if err != nil {
		panic("could not marshal object")
	}
	jsonS, _ := json.Marshal(s)
	fmt.Println(string(jsonP), jsonP)
	fmt.Println(string(jsonS), jsonS)
}
