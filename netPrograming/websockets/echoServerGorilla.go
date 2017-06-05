package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling /")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for n := 0; n < 10; n++ {
		msg := "Hello " + string(n+66)
		err = conn.WriteMessage(websocket.TextMessage, []byte(msg))

		time.Sleep(2 * time.Second)
		_, reply, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("cannot receive")
			break
		}
		fmt.Println("Received back from client:", string(reply[:]))
	}
}

func main() {
	http.HandleFunc("/", echo)
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
