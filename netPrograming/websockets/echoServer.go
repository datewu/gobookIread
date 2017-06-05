package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", websocket.Handler(echo))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func echo(ws *websocket.Conn) {
	fmt.Println("Echoing")

	for n := 0; n < 10; n++ {
		msg := "hello" + string(n+65)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			fmt.Println("cannot send")
			break
		}
		time.Sleep(2 * time.Second)
		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("cannot receive")
			break
		}
		fmt.Println("Received back from client:" + reply)
	}
}
