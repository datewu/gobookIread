package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Values RPC arguments
type Values struct {
	A, B int
}

// Quotient RPC return values
type Quotient struct {
	Que, Rem int
}

// Arith staticfy RPC methods
type Arith int

// Multiply the RPC method
func (*Arith) Multiply(args *Values, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide the other RPC method
func (*Arith) Divide(args *Values, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Que = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	/* This doesnot works
	jsonrpc.Accept(listener)

		There is no jsonrpc Accept(lis net.Listener)
	*/
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err)

	}
}
