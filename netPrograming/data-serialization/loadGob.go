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

// String statisfy fmt
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ":" + v.Address
	}
	return s
}

func main() {
	var p Person
	loadGob("person.gob", &p)
	fmt.Println("Person", p)
}

func loadGob(filename string, key interface{}) {
	inFile, err := os.Open(filename)
	checkError(err)

	encoder := gob.NewDecoder(inFile)
	err = encoder.Decode(key)
	checkError(err)
	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
