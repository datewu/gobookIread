package main

import (
	"fmt"
	"net/http"
	"sync"
)

type result struct {
	Error    error
	Response *http.Response
}

func notThatGoodErrorHandlering() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		var wg sync.WaitGroup
		responses := make(chan *http.Response)

		go func() {
			wg.Add(len(urls))
			defer close(responses)
			for _, url := range urls {
				go func(u string) {
					defer wg.Done()
					resp, err := http.Get(u)
					if err != nil {
						fmt.Println(err)
						return
					}
					select {
					case <-done:
						return
					case responses <- resp:
					}
				}(url)
			}
			wg.Wait()
		}()
		return responses
	}
	done := make(chan interface{})

	defer close(done)

	domains := []string{"https://taobao.com", "https://jd.com", "https://baidu.com", "https://fsdffd.df"}
	for response := range checkStatus(done, domains...) {
		fmt.Println("Response:", response.Status)

	}

}

func goodEoorHandlering() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan result {
		var wg sync.WaitGroup
		results := make(chan result)

		go func() {
			wg.Add(len(urls))
			defer close(results)
			for _, url := range urls {
				go func(u string) {
					defer wg.Done()
					resp, err := http.Get(u)
					r := result{err, resp}
					select {
					case <-done:
						return
					case results <- r:
					}
				}(url)
			}
			wg.Wait()
		}()
		return results
	}
	done := make(chan interface{})

	defer close(done)

	domains := []string{"https://taobao.com", "https://jd.com", "https://baidu.com", "https://fsdffd.df"}
	for result := range checkStatus(done, domains...) {
		if result.Error != nil {
			fmt.Println("error:", result.Error, "lol")
			continue
		}
		fmt.Println("Response:", result.Response.Status)
	}

}
