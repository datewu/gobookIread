package main

import (
	"fmt"
	"net"
	"os"
)

// change this to my own IP address or set 0.0.0.0
const (
	myIPAddress    = "0.0.0.0"
	ipv4HeaderSize = 20
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	localAddr, err := net.ResolveIPAddr("ip4", myIPAddress)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	remoteAddr, err := net.ResolveIPAddr("ip4", os.Args[1])
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialIP("ip:icmp", localAddr, remoteAddr)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum, fix later
	msg[3] = 0  // checksum, fix later
	msg[4] = 0  // identifier[0]
	msg[5] = 12 // identifier[1] (arbitrary)
	msg[6] = 0  // sequence[0]
	msg[7] = 38 // sequence[1] (arbitrary)
	len := 8

	// now fix checksum bytes
	check := checkSum(msg[:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	// send the message
	_, err = conn.Write(msg[:len])
	checkError(err)

	fmt.Print("Message send: ")
	for n := 0; n < 8; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()

	// receive a reply
	size, err := conn.Read(msg[:])
	checkError(err)

	fmt.Print("Message reveived: ")
	for n := ipv4HeaderSize; n < size; n++ {
		fmt.Print(" ", msg[n])
	}
	fmt.Println()
}

func checkSum(msg []byte) uint16 {
	sum := 0

	// assume even for now
	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
