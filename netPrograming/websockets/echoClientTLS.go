package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "wss://host:port")
	}
	service := os.Args[1]
	config, err := websocket.NewConfig(service, "http://localhost:8080")
	if err != nil {
		log.Fatalln("websocket dial:", err)
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	config.TlsConfig = tlsConfig

	conn, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatalln("websocket dial:", err)
	}
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				break
			}
			fmt.Println("cannot receive msg", err)
			break
		}
		fmt.Println("Received form server:", msg)
		// return the msg
		err = websocket.Message.Send(conn, msg+" from client")
		if err != nil {
			fmt.Println("Couldnot return msg")
			break
		}
	}
}
