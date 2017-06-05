package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	http.Handle("/ws", websocket.Handler(date))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func date(ws *websocket.Conn) {
	for {
		msg, _ := exec.Command("date").Output()
		err := websocket.Message.Send(ws, string(msg[:]))
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
