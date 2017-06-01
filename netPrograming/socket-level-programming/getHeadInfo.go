package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
	}
	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	chekError(err)
	log.Println(tcpAddr)
	log.Println(tcpAddr.IP)

	lo := net.ParseIP("127.0.0.1")
	tcpAddr.IP = lo
	log.Println(tcpAddr)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	chekError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	chekError(err)
	result, err := ioutil.ReadAll(conn)
	chekError(err)
	fmt.Println(string(result))
	conn.Close()
}

func chekError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

}
