package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//hash := md5.New()
	hash := hmac.New(md5.New, []byte("lol tag"))
	bytes := []byte("hello world\n")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()

	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x ", val)
	}
	fmt.Println()
	fmt.Println(hashValue)
	fmt.Println(hex.EncodeToString(hashValue))

}
