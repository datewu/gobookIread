package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Person to json
type Person struct {
	Name  Name
	Email []Email
}

// Name to json
type Name struct {
	Family, Personal string
}

// Email to json
type Email struct {
	Kind, Address string
}

func main() {
	person := Person{
		Name: Name{"Newmarch", "Jan"},
		Email: []Email{Email{"home", "jan@newmarch.name"},
			Email{"work", "j.newmarch@boxhil.edu.au"},
		}}
	saveGob("person.gob", person)
}

func saveGob(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
