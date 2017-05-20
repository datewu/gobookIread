package main

import (
	"fmt"
	"net"
	"os"
)

// DIR ... the commands
const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			conn.Close()
			return
		}
		s := string(buf[:n])
		// decode request
		if s[:2] == CD {
			chdir(conn, s[3:])
		} else if s[:3] == DIR {
			dirList(conn)
		} else if s[:3] == PWD {
			pwd(conn)
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte("ERROR"))
		return
	}
	conn.Write([]byte(s))
}

func dirList(conn net.Conn) {
	// send a blank line on termination
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}

	for _, v := range names {
		conn.Write([]byte(v + "\r\n"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
