package main

import (
	"fmt"
	"log"
	"os"
)

// LOGFILE ...
var LOGFILE = "tmp.log"

func one(aLog *log.Logger) {
	aLog.Println("-- FUNCTION one ----")
	defer aLog.Println("FUNCTION one ----")

	for index := 0; index < 10; index++ {
		aLog.Println(index)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("------ FUNCTION two")
	defer aLog.Println("FUNCTION two ------")

	for index := 0; index < 10; index++ {
		aLog.Println(index)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "logDefer ", log.LstdFlags)
	iLog.Println("Hello there")
	iLog.Println("Another log entry")

	one(iLog)
	two(iLog)
}
