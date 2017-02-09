package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("DELETE", "http://taobao.com", nil)
	res, _ := http.DefaultClient.Do(req)
	res.Body.Close()
	fmt.Println(res.Status)
}
