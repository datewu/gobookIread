package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// Args RPC arguments
type Args struct {
	A, B int
}

// Quotient RPC return values
type Quotient struct {
	Que, Rem int
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "server")
	}
	serverAddress := os.Args[1]

	client, err := rpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalln("dialing:", err)
	}

	// synchronous call
	a := Args{997, 6}

	var r int
	err = client.Call("Arith.Multiply", a, &r)
	if err != nil {
		log.Fatalln("arith multiply error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", a.A, a.B, r)

	var q Quotient
	err = client.Call("Arith.Divide", a, &q)
	if err != nil {
		log.Fatalln("arith Divide error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d \n", a.A, a.B, q.Que, q.Rem)
}
