package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	str1 := "\u04d6"
	str2 := "\u0415\u0306"

	normStr2 := norm.NFC.String(str2)
	normBytes2 := []byte(normStr2)

	bytes1 := []byte(str1)
	bytes2 := []byte(str2)

	fmt.Println("Single char", str1, "bytes", bytes1)
	fmt.Println("Composed char", str2, "bytes", bytes2)
	fmt.Println("Normalized char", normStr2, "bytes", normBytes2)

}
