package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func retry(url string) error {
	const timeout = time.Minute
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying (%d)...", err, tries)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	u := "http://jd.com"
	fmt.Println(retry(u))
}
