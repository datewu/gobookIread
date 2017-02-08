package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(15 * time.Second)
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Time out")
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) {
	for {
		data := make([]byte, 1024)
		l, err := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}

	}

}
