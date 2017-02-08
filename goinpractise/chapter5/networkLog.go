package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	// nc -lk 9999
	if err != nil {
		log.Println(err)
		panic("Failed to connect to localohst:9999")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile

	logger := log.New(conn, "example ", f)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
	logger.Fatalln("you should NOT use Fatalln, which causes defer functions NOT called")
}
