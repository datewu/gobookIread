package main

import (
	"fmt"
	"log"
)

func main() {
	x := "hello world"
	log.Printf("&x = %+v\n", &x)

	for _, x := range x {
		log.Printf("&x = %+v\n", &x)
		//x := x + 'A' - 'a'
		x = x + 'A' - 'a'
		log.Printf("&x = %+v\n", &x)
		fmt.Printf("%c", x)
	}
}
