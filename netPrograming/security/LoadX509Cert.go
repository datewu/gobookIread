package main

import (
	"crypto/x509"
	"fmt"
	"os"
)

func main() {
	certCerFile, err := os.Open("lol.dota.cert")
	checkError(err)
	derBytes := make([]byte, 1000) // bigger than the file
	count, err := certCerFile.Read(derBytes)
	checkError(err)
	certCerFile.Close()

	// trim the bytes to actual length in call
	cert, err := x509.ParseCertificate(derBytes[:count])
	checkError(err)

	fmt.Println("Name", cert.Subject.CommonName)
	fmt.Println("Not befort", cert.NotBefore)
	fmt.Println("Not after", cert.NotAfter)

}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(33)
	}
}
