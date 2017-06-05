package main

import (
	"log"
	"os"

	"golang.org/x/net/websocket"
)

type person struct {
	Name   string
	Emails []string
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "ws://host:port")
	}
	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "http://localhost:8080")

	if err != nil {
		log.Fatalln("websocket dial:", err)
	}
	p := person{"Jan", []string{"ja@n.com", "lol@dota.com"}}

	err = websocket.JSON.Send(conn, p)
	if err != nil {
		log.Fatalln("websocket JSON send failed:", err)
	}
}
