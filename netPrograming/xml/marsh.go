package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type person struct {
	XMLName name    `xml:"person"`
	Name    name    `xml:"name"`
	Email   []email `xml:"email"`
}

type name struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}

type email struct {
	Type    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func main() {

	p := person{
		Name: name{"Newmarch", "Jan"},
		Email: []email{email{"home", "jan"},
			email{"work", "lol"}}}

	buff, err := xml.Marshal(p)
	if err != nil {
		log.Fatalln("Fatal error", err)
	}
	fmt.Println(string(buff))
}
