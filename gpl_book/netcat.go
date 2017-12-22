package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//	netcatReceiver()
	netcat()
}

func netcat() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	doneStream := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		doneStream <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-doneStream
}

func netcatReceiver() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
