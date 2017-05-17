package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr \n", os.Args[0])
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	mask := addr.DefaultMask()
	ones, bits := mask.Size()

	network := addr.Mask(mask)
	fmt.Println("Address is ", addr.String(),
		"\n Mask length is ", bits,
		"\n Leading ones count is ", ones,
		"\n Mask is (hex) ", mask.String(),
		"\n Network is ", network.String())

}
