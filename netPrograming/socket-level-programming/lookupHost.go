package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
	c, _ := net.LookupCNAME(name)
	fmt.Println("cname: ", c)
}
