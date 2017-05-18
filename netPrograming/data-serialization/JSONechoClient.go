package main

import (
	"encoding/json"
	"fmt"
	"net"
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

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ":" + v.Address
	}
	return s
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	person := Person{
		Name: Name{"Newmarch", "Jan"},
		Email: []Email{Email{"home", "jan@newmarch.name"},
			Email{"work", "j.newmarch@boxhil.edu.au"},
		}}

	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
