package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
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
	rpc.HandleHTTP()

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
