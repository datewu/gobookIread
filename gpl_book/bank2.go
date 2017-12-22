package main

import (
	"fmt"
)

var (
	sema = make(chan struct{}, 1)
	b    int
)

func deposit(amount int) {
	sema <- struct{}{}
	b += amount
	<-sema
}

func balance() int {
	sema <- struct{}{}
	r := b
	<-sema
	return r

}

func main() {

	fmt.Println(balance())
	deposit(100)
	fmt.Println("After deposit 100:", balance())
}
