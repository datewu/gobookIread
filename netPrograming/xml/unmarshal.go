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

	str := `<?xml version="1.0" encoding="utf-8"?>
<person>
 <name>
 <family> Newmarch </family>
 <personal> Jan </personal>
 </name>
 <email type="personal">
 jan@newmarch.name
 </email>
 <email type="work">
 j.newmarch@boxhill.edu.au
 </email>
</person>`

	var p person

	err := xml.Unmarshal([]byte(str), &p)
	if err != nil {
		log.Fatalln("Fatal error", err)
	}
	fmt.Println("Family name: \"" + p.Name.Family + "\"")
	fmt.Println("Second email address: \"" + p.Email[1].Address + "\"")

}
