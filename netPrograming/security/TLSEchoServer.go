package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("lol.dota.pem", "private.pem")
	checkError(err)

	config := tls.Config{Certificates: []tls.Certificate{cert}}
	now := time.Now()
	config.Time = func() time.Time {
		return now
	}
	config.Rand = rand.Reader

	listener, err := tls.Listen("tcp", ":1200", &config)
	checkError(err)
	fmt.Println("Listening on :1200")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Accept")
		go handleClient(conn)
	}
}

func handleClient(c net.Conn) {
	defer c.Close()
	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := c.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("receive from client", string(buf[:n]))
		_, err = c.Write(buf[:n])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err)
	}
}
