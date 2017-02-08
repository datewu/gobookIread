package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// error variables
var (
	ErrTimeout  = errors.New("the request timed out")
	ErrRejected = errors.New("the request was rejected")
)

var random = rand.New(rand.NewSource(35))

func main() {
	response, err := sendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		response, err = sendRequest("Hello")
	}
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response)
	}
}

func sendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}

}
