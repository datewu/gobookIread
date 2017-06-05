package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "ws://host:port")
	}
	service := os.Args[1]

	header := make(http.Header)
	header.Add("Origin", "http://localhost:8080")

	conn, _, err := websocket.DefaultDialer.Dial(service, header)
	if err != nil {
		log.Fatalln("websocket dial:", err)
	}
	for {
		_, reply, err := conn.ReadMessage()
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				fmt.Println("EOF from server")
				break
			}
			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				fmt.Println("close from server")
				break

			}
			fmt.Println("cannot receive msg", err)
			break
		}
		fmt.Println("Received form server:", string(reply[:]))
		// return the msg
		reply = append(reply, byte('L'))
		err = conn.WriteMessage(websocket.TextMessage, reply)
		if err != nil {
			fmt.Println("Couldnot return msg")
			break
		}
	}
}
