package main

import (
	"flag"
	"os"
	"strings"
)

func main() {

}
func echo1() string {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	//fmt.Println(s)
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, a := range os.Args[1:] {
		s += sep + a
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

var sep = flag.String("sep", " ", "seperateor, default is ' '")
var n = flag.Bool("n", false, "omit trailing newline")

func echo4() string {
	flag.Parse()
	s := strings.Join(os.Args[1:], *sep)
	if !*n {
		s += "\n"
	}
	return s
}
