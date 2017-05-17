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
		log.Fatalln("Usage: ", os.Args[0], "hostname")
	}
	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		log.Fatalln("Resolution error", err.Error())
	}
	fmt.Println("Resolved address is ", addr.String())
}
