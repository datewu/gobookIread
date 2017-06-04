package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usgae:", os.Args[0], "file")
	}

	file := os.Args[1]
	bytes, err := ioutil.ReadFile(file)
	checkError(err)

	r := strings.NewReader(string(bytes))

	parser := xml.NewDecoder(r)

	var depth int
	for {
		token, err := parser.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
			depth++

		case xml.EndElement:
			depth--
			elmt := xml.EndElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
		case xml.CharData:
			bytes := xml.CharData(t)
			printElmt("\""+string([]byte(bytes))+"\"", depth)
		case xml.Comment:
			printElmt("Comment", depth)
		case xml.ProcInst:
			printElmt("ProcInst", depth)
		case xml.Directive:
			printElmt("Directive", depth)
		default:
			fmt.Println("Unknown")
		}

	}

}

func printElmt(s string, d int) {
	for n := 0; n < d; n++ {
		fmt.Print(" ")
	}
	fmt.Println(s)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err)
	}

}
