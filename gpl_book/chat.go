package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}

}

type clientStream chan<- string // an outgoing message channel

var (
	enteringStream = make(chan clientStream)
	leavingStream  = make(chan clientStream)
	messageStream  = make(chan string)
)

func broadcaster() {
	clients := make(map[clientStream]bool)

	for {
		select {
		case msg := <-messageStream:
			for cliStream := range clients {
				cliStream <- msg
			}
		case cliStream := <-enteringStream:
			clients[cliStream] = true
		case cliStream := <-leavingStream:
			delete(clients, cliStream)
			close(cliStream)
		}
	}
}

func handleConn(conn net.Conn) {
	cliStream := make(chan string)
	go clientWrite(conn, cliStream)

	who := conn.RemoteAddr().String()
	cliStream <- "Your ID is: " + who

	messageStream <- "Hi everyone, " + who + " has arrived"
	enteringStream <- cliStream

	cliInput := bufio.NewScanner(conn)
	for cliInput.Scan() {
		messageStream <- who + ": " + cliInput.Text()
	}
	// NOTE: ignoring potential errors from cliInput.Err()

	leavingStream <- cliStream
	messageStream <- who + " has just left!"
	conn.Close()
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
