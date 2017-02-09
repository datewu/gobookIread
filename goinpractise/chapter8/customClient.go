package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	cc :=
		&http.Client{Timeout: time.Second}
	res, err := cc.Get("http://taobao.com")
	if err != nil {
		log.Println(err)
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
