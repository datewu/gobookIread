package main

import (
	"log"
	"time"
)

func main() {
	bigSlow()
}

func bigSlow() {
	defer trace("Tracing")()
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Println("enter", msg)
	return func() {
		log.Println("exit", msg, ": After", time.Since(start))
	}
}
