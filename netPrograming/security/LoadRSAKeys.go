package main

import (
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	var key rsa.PrivateKey
	loadKey("private.key", &key)

	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	var publickey rsa.PublicKey
	loadKey("public.key", &publickey)
	fmt.Println("Public key modulus", publickey.N.String())
	fmt.Println("Public key exponent", publickey.E)
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
