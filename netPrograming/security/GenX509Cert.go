package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader
	var key rsa.PrivateKey
	loadKey("private.key", &key)

	now := time.Now()
	//then := now.Add(60 *60 * 24 *365 *100*1000*1000) // one year
	then := now.AddDate(1, 0, 0)

	tpl := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "lol.dota",
			Organization: []string{"big dream"},
		},
		NotBefore: now,
		NotAfter:  then,

		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:     true,
		DNSNames: []string{"lol.dota", "localhost"},
	}

	derBytes, err := x509.CreateCertificate(random, &tpl, &tpl, &key.PublicKey, &key)
	checkError(err)

	certCerFilie, err := os.Create("lol.dota.cert")
	checkError(err)
	certCerFilie.Write(derBytes)
	certCerFilie.Close()

	certPEMFile, err := os.Create("lol.dota.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certPEMFile.Close()

	keyPEMFile, err := os.Create("private.pem.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(&key)})

	keyPEMFile.Close()

}

func loadKey(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)

	inFile.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(33)
	}
}
