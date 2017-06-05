package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type person struct {
	Name   string
	Emails []string
}

func receivePerson(ws *websocket.Conn) {
	var p person
	err := websocket.JSON.Receive(ws, &p)
	if err != nil {
		fmt.Println("Cannot receive")
	} else {
		fmt.Println("Name:", p.Name)
		for _, v := range p.Emails {
			fmt.Println("An eamil:", v)
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(receivePerson))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
