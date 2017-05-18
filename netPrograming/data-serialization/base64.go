package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	eithtBitData := []byte{1, 10, 100, 200, 11, 101, 107, 208}

	enc := base64.StdEncoding.EncodeToString(eithtBitData)
	dec, _ := base64.StdEncoding.DecodeString(enc)

	fmt.Println("Original data ", eithtBitData, string(eithtBitData), len(string(eithtBitData)))
	fmt.Println("Encoded string ", enc)
	fmt.Println("Decoded data ", dec, string(dec))
}
