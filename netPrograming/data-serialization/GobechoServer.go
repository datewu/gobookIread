package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

// Person to gob
type Person struct {
	Name  Name
	Email []Email
}

// Name to gob
type Name struct {
	Family, Personal string
}

// Email to gob
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

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var p Person
			decoder.Decode(&p)
			fmt.Println(p)
			p.Name.Family = "lol"
			encoder.Encode(p)
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
