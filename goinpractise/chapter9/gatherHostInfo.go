package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(name)
	addres, err := net.LookupHost(name)
	if err != nil {
		log.Println(err)
		return
	}

	for _, a := range addres {
		fmt.Println(a)
	}

}
