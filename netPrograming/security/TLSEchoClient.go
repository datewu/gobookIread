package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "host:port")
	}
	service := os.Args[1]

	// Load the PEM self-signed sertificate
	certPemFile, err := os.Open("lol.dota.pem")
	checkError(err)

	pemBytes := make([]byte, 1000) // bigger than the file
	_, err = certPemFile.Read(pemBytes)
	checkError(err)
	certPemFile.Close()

	// Create a new certificat pool
	certPool := x509.NewCertPool()
	// add our custom certifucate
	ok := certPool.AppendCertsFromPEM(pemBytes)
	if !ok {
		log.Fatalln("PEM read failed")
	} else {
		fmt.Println("PEM read OK, yah")
	}

	// Dial, using a config with root cert seet to ours
	conn, err := tls.Dial("tcp", service, &tls.Config{RootCAs: certPool})
	checkError(err)

	// Now write and read
	for n := 0; n < 10; n++ {
		fmt.Println("Writting to server..", string(n+66))
		conn.Write([]byte("Hello " + string(n+66)))

		var buf [512]byte
		n, err := conn.Read(buf[:])
		checkError(err)
		fmt.Println("Response from server:", string(buf[:n]))
	}
	conn.Close()
}

func handleClient(c net.Conn) {
	defer c.Close()
	var buf [512]byte
	for {
		fmt.Println("Trying to read")
		n, err := c.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = c.Write(buf[:n])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err)
	}
}
