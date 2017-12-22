package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	b  int
)

func deposit(amount int) {
	mu.Lock()
	b += amount
	mu.Unlock()
}

func balance() int {
	mu.Lock()
	defer mu.Unlock()
	return b

}

func main() {

	fmt.Println(balance())
	deposit(100)
	fmt.Println("After deposit 100:", balance())
}
