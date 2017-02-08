package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(conf.Path)
}
