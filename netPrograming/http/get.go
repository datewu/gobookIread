package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host:port")
		os.Exit(2)
	}
	url := os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(5)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))

	contentTypes := response.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannot handle", contentTypes)
		os.Exit(5)
	}
	fmt.Println("The response body is")
	time.Sleep(time.Second * 15)
	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Println(string(buf[:n]))
	}
}

func acceptableCharset(contentTypes []string) bool {
	for _, cType := range contentTypes {
		if strings.Index(cType, "utf-8") != -1 {
			return true
		}
	}
	return false

}
