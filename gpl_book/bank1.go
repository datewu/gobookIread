package main

import (
	"fmt"
	"log"
)

var (
	depositStream  = make(chan int)
	withdrawStream = make(chan ws)
	balanceStream  = make(chan int)
)

type ws struct {
	amount   int
	callback chan bool
}

func deposit(amount int) {
	depositStream <- amount
}

func balance() int {
	return <-balanceStream
}

func withdraw(amount int) (b bool) {
	w := ws{amount, make(chan bool)}
	withdrawStream <- w
	b = <-w.callback
	log.Println(b)
	return
}

func teller() {
	var balance int // balance is confied to teller goroutine
	for {
		select {
		case amount := <-depositStream:
			balance += amount
		case balanceStream <- balance:
			//log.Println("he is cheking his balance")
		case w := <-withdrawStream:
			if w.amount <= balance {
				balance -= w.amount
				w.callback <- true
			} else {
				w.callback <- false
			}
		}
	}
}

func main() {
	go teller()

	fmt.Println(balance())
	deposit(100)
	fmt.Println("After deposit 100:", balance())
	withdraw(200)
	fmt.Println("After withdraw 200:", balance())
	withdraw(20)
	fmt.Println("After withdraw 20:", balance())
}
