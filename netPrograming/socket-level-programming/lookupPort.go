package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type service\n",
			os.Args[0])
		os.Exit(1)
	}
	networkType, service := os.Args[1], os.Args[2]

	port, err := net.LookupPort(networkType, service)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}
	fmt.Println("Service port ", port)
}
